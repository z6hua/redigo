package server

import (
	"github.com/satori/go.uuid"
	"redigo/redigo/constants"
	"redigo/redigo/constants/enums"
	"redigo/redigo/db"
	"sync"
)

type Server struct {
	runId      string
	flag       enums.ServerType
	master     *Master
	slaves     map[string]Slave
	replBuffer *ReplBuffer
	replOffset int
	db         db.DB
	psyncLock  sync.Mutex
	psyncFlag  string
}

var global = Server{
	runId:      uuid.NewV4().String(),
	flag:       enums.Master,
	slaves:     make(map[string]Slave),
	replBuffer: NewReplBuffer(constants.DefaultReplBufferSize),
	replOffset: 0,
	db: db.DB{
		Ht: make(map[string]string),
	},
}

var Funcs = map[string]any{
	"Set": DB().Set,
	"Del": DB().Del,
}

func GetRunId() string {
	return global.runId
}

func DB() *db.DB {
	return &(global.db)
}

func Slaves() *map[string]Slave {
	return &(global.slaves)
}

func GetMaster() *Master {
	return global.master
}

func SetMaster(m *Master) {
	global.master = m
}

func GetType() enums.ServerType {
	return global.flag
}

func SetType(t enums.ServerType) {
	global.flag = t
}

func GetReplInfo() (masterRunId string, replOffset int) {
	masterRunId = global.master.RunId
	replOffset = global.replOffset
	return
}

func GetReplOffset() int {
	return global.replOffset
}

func SetReplOffset(offset int) {
	global.replOffset = offset
}

func GetReplBuffer() *ReplBuffer {
	return global.replBuffer
}

func PsyncLock() {
	global.psyncLock.Lock()
}

func PsyncUnlock() {
	global.psyncLock.Unlock()
}

func GetPsyncFlag() string {
	return global.psyncFlag
}

func SetPsyncFlag(s string) {
	global.psyncFlag = s
}
