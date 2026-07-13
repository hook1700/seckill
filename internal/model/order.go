package model

type Order struct {
	UserID     int64 `json:"user_id"`
	ActivityID int64 `json:"activity_id"`
	OrderID    int64 `json:"order_id"`
}
