package routers

import (
	"github.com/gin-gonic/gin"
	"redigo/redigo/controllers/command"
)

func InitCommandRouters(e *gin.Engine) {
	e.GET("/getall", command.GetAllCommand)
	e.GET("/get", command.GetCommand)
	e.GET("/set", command.SetCommand)
	e.GET("/del", command.DelCommand)
}
