package repo

type Order struct {
	ID         int64  `gorm:"primaryKey;autoIncrement"`
	OrderID    string `gorm:"type:varchar(64);uniqueIndex"`
	UserID     int64  `gorm:"not null"`
	ActivityID int64  `gorm:"not null"`
	Status     int8   `gorm:"default:1"`
}

const (
	SeckillOK       = 1
	SeckillSoldOut  = -1
	SeckillRepeated = -2
)
