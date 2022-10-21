package repl

import (
	"github.com/gin-gonic/gin"
	"log"
	"net"
	"net/http"
	"redigo/redigo/constants/enums"
	"redigo/redigo/logics"
	"redigo/redigo/server"
	"redigo/redigo/utils"
	"strings"
)

func SlaveList(c *gin.Context) {
	c.JSON(http.StatusOK, server.Slaves())
}

// SlaveOf 请求成为其他服务器的从服务器（认大哥）
func SlaveOf(c *gin.Context) {
	host := c.Query("host")
	port := c.Query("port")

	url := utils.GenURL(host, port, "/slaveof_r")
	params := map[string]any{
		"runId": server.GetRunId(),
	}

	resp, err := utils.RequestGet(url, params)
	if err != nil {
		c.String(http.StatusNotFound, "cannot connected master server")
		log.Panic(err)
		return
	}

	master := &server.Master{}
	utils.ReadResponse2Obj(resp, master)
	master.Host = host
	master.Port = port
	server.SetMaster(master)
	server.SetType(enums.Slave)
	server.SetReplOffset(-1)
	go logics.HearBeatPsync()
	c.JSON(http.StatusOK, "ok")
}

// SlaveOfR 处理从服务器的申请请求（处理任大哥的请求）
func SlaveOfR(c *gin.Context) {
	slaveRunId := c.Query("runId")
	host, port, _ := net.SplitHostPort(strings.TrimSpace(c.Request.RemoteAddr))
	if host == "::1" {
		host = "localhost"
	}
	slaves := *(server.Slaves())
	slaves[slaveRunId] = server.Slave{RunId: slaveRunId, Host: host, Port: port}
	c.JSON(http.StatusOK, gin.H{
		"runId": server.GetRunId(),
	})
}

// Psync 从服务器同步主服务器数据库状态
func Psync(c *gin.Context) {
	master := server.GetMaster()
	if server.GetType() != enums.Slave || master == nil {
		c.String(http.StatusOK, "current server not a slave server")
		return
	}
	logics.Psync()
}

// PsyncR 主服务器对从服务器同步数据库状态的处理
func PsyncR(c *gin.Context) {
	slaveRunId := c.Query("runId")
	masterRunId := c.Query("masterRunId")
	offset := utils.Int(c.Query("offset"))
	slaves := *(server.Slaves())
	if _, ok := slaves[slaveRunId]; !ok {
		c.JSON(http.StatusOK, gin.H{
			"message": "server not my slave server",
		})
		return
	}
	c.JSON(http.StatusOK, logics.PsyncMaster(masterRunId, offset))
}
