package repo

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitMySQL(dsn string, maxOpen, maxIdle int) {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("mysql connect failed: %v", err))
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(maxOpen)
	sqlDB.SetMaxIdleConns(maxIdle)
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
