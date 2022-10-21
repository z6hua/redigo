package logics

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"redigo/redigo/constants/enums"
	rpc "redigo/redigo/rpc"
	"redigo/redigo/server"
	"strings"
	"time"
)

func HearBeatPsync() {
	psyncFlag := uuid.NewV4().String()
	server.SetPsyncFlag(psyncFlag)
	for psyncFlag == server.GetPsyncFlag() {
		Psync()
		time.Sleep(time.Second * 10)
	}
}

// Psync 从服务器向主服务器PSYNC同步处理逻辑
func Psync() {
	server.PsyncLock()
	defer server.PsyncUnlock()
	masterRunId, offset := server.GetReplInfo()
	runId := server.GetRunId()
	replRPC := rpc.NewReplRPC()
	data := replRPC.Psync(masterRunId, runId, offset)
	server.SetReplOffset(int(data["offset"].(float64)))
	server.GetMaster().RunId = data["runId"].(string)
	syncType := enums.ReplSYNCType(data["resync"].(string))
	switch syncType {
	case enums.FullReSync:
		dbData := data["db_data"].(map[string]any)
		dbHt := make(map[string]string)
		for k, v := range dbData {
			dbHt[k] = fmt.Sprintf("%v", v)
		}
		FullReSync(dbHt)
	case enums.PartReSync:
		dbData := data["db_data"].(string)
		PartReSync(dbData)
	}
}

// PsyncMaster 主服务器处理从服务器PSYNC同步逻辑
func PsyncMaster(runId string, offset int) gin.H {
	res := gin.H{}
	res["offset"] = server.GetReplOffset()
	res["runId"] = server.GetRunId()
	isFullReSync := offset == -1 || runId != server.GetRunId() || server.GetReplOffset()-offset > server.GetReplBuffer().Size()
	if isFullReSync {
		res["resync"] = enums.FullReSync
		res["db_data"] = server.DB().Ht
	} else {
		res["resync"] = enums.PartReSync
		if offset == server.GetReplOffset() {
			res["db_data"] = ""
		} else {
			offset = server.GetReplBuffer().Len() - server.GetReplOffset() + offset
			res["db_data"] = string(server.GetReplBuffer().Get(offset))
		}
	}
	return res
}

func FullReSync(data map[string]string) {
	server.DB().Ht = data
}

func PartReSync(data string) {
	commandSlice := strings.Split(data, "\n")
	for _, commandStr := range commandSlice {
		command := strings.Split(commandStr, " ")
		commandFunc := command[0]
		commandParams := command[1:]
		commandParamsSlice := make([]any, len(commandParams))
		for i, val := range commandParams {
			commandParamsSlice[i] = val
		}
		switch commandFunc {
		case "Set":
			server.DB().Set(commandParams[0], commandParams[1])
		case "Del":
			server.DB().Del(commandParams[0])
		}
	}
}
