package mq

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"seckill/internal/repo"

	"github.com/segmentio/kafka-go"
)

func StartConsumer(brokers, topic, groupID string) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokers},
		Topic:   topic,
		GroupID: groupID,
	})

	log.Println("kafka consumer started")

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("kafka read error: %v", err)
			continue
		}

		var event repo.OrderEvent
		if err := json.Unmarshal(msg.Value, &event); err != nil {
			log.Printf("unmarshal error: %v", err)
			continue
		}

		userID, _ := strconv.ParseInt(event.UserID, 10, 64)
		activityID, _ := strconv.ParseInt(event.ActivityID, 10, 64)

		if err := repo.SaveOrder(userID, activityID); err != nil {
			log.Printf("save order failed: %v", err)
		} else {
			log.Printf("order saved: user=%s activity=%s", event.UserID, event.ActivityID)
		}
	}
}
