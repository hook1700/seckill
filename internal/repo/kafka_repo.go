package repo

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

var kafkaWriter *kafka.Writer

func InitKafka(brokers string) {
	kafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(brokers),
		Balancer: &kafka.LeastBytes{},
	}
}

type OrderEvent struct {
	UserID     string `json:"user_id"`
	ActivityID string `json:"activity_id"`
}

func SendOrderEvent(ctx context.Context, userID, activityID string) error {
	data, err := json.Marshal(OrderEvent{
		UserID:     userID,
		ActivityID: activityID,
	})
	if err != nil {
		return fmt.Errorf("marshal order event failed: %w", err)
	}
	return kafkaWriter.WriteMessages(ctx, kafka.Message{Value: data})
}
