package routers

import "github.com/gin-gonic/gin"

func InitRouters(engine *gin.Engine) {
	InitCommandRouters(engine)
	InitReplRouters(engine)
}
