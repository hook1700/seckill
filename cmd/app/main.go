package main

import (
	"log"
	"seckill/internal/api"
	"seckill/internal/config"
	"seckill/internal/repo"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. 加载配置
	cfg := config.Load()

	// 2. 初始化依赖
	repo.InitRedis(cfg.RedisAddr)
	//repo.InitKafka(cfg.KafkaBrokers)
	//repo.InitMySQL(cfg.MySQLDSN)

	// 3. 启动 Gin（压测时切 release）
	r := gin.Default()
	api.RegisterRouter(r)

	// 4. 启动 Kafka 消费（异步落单）
	//go api.StartOrderConsumer()

	log.Fatal(r.Run(":8080"))
}
