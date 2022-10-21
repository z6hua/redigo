package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"redigo/redigo/middlewares"
	"redigo/redigo/routers"
)

var port string

func init() {
	flag.StringVar(&port, "port", "8080", "listening port")
}

func main() {
	// 启动命令 go run main.go -port 8080
	flag.Parse()

	app := gin.Default()
	app.Use(middlewares.CommandMiddleware())
	routers.InitRouters(app)
	app.Run(":" + port)
}
