package main

import (
	"fmt"
	"log"
	"runtime"
	"seckill/config"
	"seckill/internal/mq"
	"seckill/internal/repo"

	"seckill/internal/api"
)

func main() {
	cfg := config.Load()

	if cfg.App.GOMAXPROCS > 0 {
		runtime.GOMAXPROCS(cfg.App.GOMAXPROCS)
	}

	repo.InitRedis(cfg.Redis.Addr, cfg.Redis.PoolSize)
	repo.InitMySQL(cfg.MySQL.DSN, cfg.MySQL.MaxOpenConns, cfg.MySQL.MaxIdleConns)
	repo.InitKafka(cfg.Kafka.Brokers)

	go mq.StartConsumer(cfg.Kafka.Brokers, cfg.Kafka.Topic, cfg.Kafka.ConsumerGroup)

	r := api.RegisterRouter()
	log.Printf("seckill server start at :%d", cfg.App.Port)
	log.Fatal(r.Run(fmt.Sprintf(":%d", cfg.App.Port)))
}
