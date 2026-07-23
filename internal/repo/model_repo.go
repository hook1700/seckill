package repo

type Order struct {
	ID         int64  `gorm:"primaryKey;autoIncrement"`
	OrderID    string `gorm:"type:varchar(64);uniqueIndex"`
	UserID     int64  `gorm:"not null"`
	ActivityID int64  `gorm:"not null"`
	Status     int8   `gorm:"default:1"`
}

// ✅ 关键：显式指定表名
func (Order) TableName() string {
	return "seckill_order"
}

const (
	SeckillOK       = 1
	SeckillSoldOut  = -1
	SeckillRepeated = -2
)
