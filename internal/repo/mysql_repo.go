package repo

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitMySQL(dsn string, maxOpen, maxIdle int) {
	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			sqlDB, _ := db.DB()
			sqlDB.SetMaxOpenConns(maxOpen)
			sqlDB.SetMaxIdleConns(maxIdle)
			log.Println("mysql connected")
			return
		}
		log.Printf("mysql not ready, retry %d: %v", i+1, err)
		time.Sleep(2 * time.Second)
	}
	panic("mysql connect failed after retries")
}

type SeckillOrder struct {
	ID         int64 `gorm:"primaryKey;autoIncrement"`
	UserID     int64 `gorm:"not null;uniqueIndex:uk_user_activity"`
	ActivityID int64 `gorm:"not null;uniqueIndex:uk_user_activity"`
	Status     int8  `gorm:"default:1"`
}

func SaveOrder(userID, activityID int64) error {
	return db.Create(&SeckillOrder{
		UserID:     userID,
		ActivityID: activityID,
	}).Error
}
