package main

import (
	"log"
	"runtime"

	"seckill/internal/config"
	"seckill/internal/repo"
	"seckill/internal/worker"

	"seckill/internal/api"
)

func main() {
	cfg := config.Load()

	if cfg.App.GOMAXPROCS > 0 {
		runtime.GOMAXPROCS(cfg.App.GOMAXPROCS)
	}

	repo.InitRedis(cfg.Redis.Addr, cfg.Redis.PoolSize)
	repo.InitMySQL(cfg.MySQL.DSN, cfg.MySQL.MaxOpenConns, cfg.MySQL.MaxIdleConns)

	go worker.StartOrderWorker("1")

	r := api.RegisterRouter()
	log.Printf("seckill server start at :%d", cfg.App.Port)
	log.Fatal(r.Run(":8080"))
}
