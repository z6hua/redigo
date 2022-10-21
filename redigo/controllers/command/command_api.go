package command

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"redigo/redigo/server"
)

func GetAllCommand(c *gin.Context) {
	c.JSON(http.StatusOK, server.DB().GetAll())
}

func GetCommand(c *gin.Context) {
	key := c.Query("key")
	c.String(http.StatusOK, server.DB().Get(key))
}

func SetCommand(c *gin.Context) {
	key := c.Query("key")
	val := c.Query("val")
	c.String(http.StatusOK, server.DB().Set(key, val))
}

func DelCommand(c *gin.Context) {
	key := c.Query("key")
	c.String(http.StatusOK, server.DB().Del(key))
}
