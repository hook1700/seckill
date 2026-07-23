package repo

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func SaveOrder(userID, activityID int64, orderID string) error {
	return db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "activity_id"}},
		DoNothing: true,
	}).Create(&Order{
		OrderID:    orderID,
		UserID:     userID,
		ActivityID: activityID,
	}).Error
}
