package main

import (
	"os"
	"seckill/internal/api"
	"seckill/internal/repo"

	"github.com/gin-gonic/gin"
)

func main() {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "redis:6379" // 兜底
	}

	repo.InitRedis(redisAddr)
	r := gin.Default()
	api.RegisterRouter(r)
	r.Run(":8080")
}
