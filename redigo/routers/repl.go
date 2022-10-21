package routers

import (
	"github.com/gin-gonic/gin"
	"redigo/redigo/controllers/repl"
)

func InitReplRouters(e *gin.Engine) {
	e.GET("/slave_list", repl.SlaveList)

	e.GET("/slaveof", repl.SlaveOf)
	e.GET("/slaveof_r", repl.SlaveOfR)

	e.GET("/psync", repl.Psync)
	e.GET("/psync_r", repl.PsyncR)
}
