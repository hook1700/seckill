package api

import (
	"net/http"
	"seckill/internal/service"

	"github.com/gin-gonic/gin"
)

func SeckillHandler(c *gin.Context) {
	userID := c.Query("user_id")
	activityID := c.Query("activity_id")

	if userID == "" || activityID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	result, err := service.DoSeckill(c.Request.Context(), userID, activityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result == -1 {
		c.JSON(http.StatusOK, gin.H{"msg": "sold out"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "queued", "order_id": result})
}
