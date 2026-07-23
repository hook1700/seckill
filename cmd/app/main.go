package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"runtime"
	"time"

	"seckill/internal/config"
	"seckill/internal/repo"
	"seckill/internal/worker"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"seckill/internal/api"
)

func main() {
	cfg := config.Load()

	if cfg.App.GOMAXPROCS > 0 {
		runtime.GOMAXPROCS(cfg.App.GOMAXPROCS)
	}

	waitForMySQL(cfg.MySQL.DSN, 30)
	waitForRedis(cfg.Redis.Addr, 30)

	repo.InitRedis(cfg.Redis.Addr, cfg.Redis.PoolSize)
	repo.InitMySQL(cfg.MySQL.DSN, cfg.MySQL.MaxOpenConns, cfg.MySQL.MaxIdleConns)

	go worker.StartOrderWorker("1")

	r := api.RegisterRouter()
	log.Printf("seckill server start at :%d", cfg.App.Port)
	log.Fatal(r.Run(":8080"))
}

func waitForMySQL(dsn string, maxRetry int) {
	var db *gorm.DB
	var err error

	for i := 1; i <= maxRetry; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, err := db.DB()
			if err == nil && sqlDB.Ping() == nil {
				log.Println("✅ mysql is ready")
				return
			}
		}
		log.Printf("⏳ mysql not ready, retry %d/%d: %v", i, maxRetry, err)
		time.Sleep(2 * time.Second)
	}
	log.Fatal("❌ mysql not ready after retries")
}

func waitForRedis(addr string, maxRetry int) {
	ctx := context.Background()

	for i := 1; i <= maxRetry; i++ {
		rdb := redis.NewClient(&redis.Options{
			Addr: addr,
		})
		if err := rdb.Ping(ctx).Err(); err == nil {
			log.Println("✅ redis is ready")
			rdb.Close()
			return
		} else {
			log.Printf("⏳ redis not ready, retry %d/%d: %v", i, maxRetry, err)
		}
		rdb.Close()
		time.Sleep(2 * time.Second)
	}
	log.Fatal("❌ redis not ready after retries")
}
