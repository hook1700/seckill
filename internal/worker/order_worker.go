package worker

import (
	"context"
	"log"
	"strconv"
	"time"

	"seckill/internal/repo"
)

func StartOrderWorker(activityID string) {
	for {
		userIDStr, orderID, ok := repo.PopOrder(context.Background(), activityID)
		if !ok {
			time.Sleep(100 * time.Millisecond)
			continue
		}

		uID, _ := strconv.ParseInt(userIDStr, 10, 64)
		if err := repo.SaveOrder(uID, 1, orderID); err != nil {
			log.Printf("save order failed: %v, orderID=%s", err, orderID)
			continue
		}
		log.Printf("order saved: orderID=%s", orderID)
	}
}
