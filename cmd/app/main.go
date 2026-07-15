package main

import (
	"log"
	"seckill/internal/api"
	"seckill/internal/repo"

	"github.com/gin-gonic/gin"
)

func main() {
	repo.InitRedis("localhost:6379")

	r := gin.Default()
	api.RegisterRouter(r)

	log.Fatal(r.Run(":8080"))
}
