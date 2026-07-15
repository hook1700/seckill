package api

import "github.com/gin-gonic/gin"

func RegisterRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/seckill", SeckillHandler)
	return r
}
