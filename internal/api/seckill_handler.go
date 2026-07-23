package api

import (
	"seckill/internal/service"

	"github.com/gin-gonic/gin"
)

func SeckillHandler(c *gin.Context) {
	userID := c.Query("user_id")
	activityID := c.Query("activity_id")

	orderID, err := service.DoSeckill(c.Request.Context(), userID, activityID)
	if err != nil {
		if err == service.ErrSoldOut {
			c.JSON(200, gin.H{"msg": "sold out"})
		} else {
			c.JSON(500, gin.H{"error": "internal"})
		}
		return
	}
	c.JSON(200, gin.H{"msg": "queued", "order_id": orderID})
}
