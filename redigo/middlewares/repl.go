package middlewares

import (
	"github.com/gin-gonic/gin"
	"redigo/redigo/server"
	"redigo/redigo/utils"
	"strings"
)

var listeningCommand = map[string]bool{
	"/set": true,
	"/del": true,
}

func CommandMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		uri := strings.Split(c.Request.RequestURI, "?")[0]
		if _, ok := listeningCommand[uri]; ok {
			var command string
			switch uri {
			case "/set":
				command = utils.GenCommand("Set", c.Query("key"), c.Query("val"))
			case "/del":
				command = utils.GenCommand("Del", c.Query("key"))
			}
			handleReplBuffer(command)
		}
	}
}

func handleReplBuffer(command string) {
	offset := server.GetReplOffset()
	server.SetReplOffset(offset + len([]byte(command)))
	server.GetReplBuffer().Push(command)
}
