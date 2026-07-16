package mq

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"seckill/internal/repo"

	"github.com/segmentio/kafka-go"
)

func StartConsumer(brokers, topic, groupID string) {

	waitForKafka(brokers)
	// 等待 Kafka 可用
	var reader *kafka.Reader
	for i := 1; i <= 30; i++ {
		conn, err := kafka.Dial("tcp", brokers)
		if err == nil {
			conn.Close()
			log.Println("✅ kafka is ready")
			break
		}
		log.Printf("⏳ kafka not ready, retry %d/30: %v", i, err)
		time.Sleep(2 * time.Second)
	}

	reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{brokers},
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 1,
		MaxBytes: 10e6,
	})

	log.Println("✅ kafka consumer started")

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("❌ kafka read error: %v", err)
			time.Sleep(1 * time.Second)
			continue
		}

		var event repo.OrderEvent
		if err := json.Unmarshal(m.Value, &event); err != nil {
			log.Printf("❌ unmarshal error: %v", err)
			continue
		}

		userID, _ := strconv.ParseInt(event.UserID, 10, 64)
		activityID, _ := strconv.ParseInt(event.ActivityID, 10, 64)

		if err := repo.SaveOrder(userID, activityID); err != nil {
			log.Printf("❌ save order failed: %v", err)
		} else {
			log.Printf("✅ order saved: user=%s activity=%s", event.UserID, event.ActivityID)
		}
	}
}

func waitForKafka(brokers string) {
	for i := 1; i <= 40; i++ {
		conn, err := kafka.Dial("tcp", brokers)
		if err == nil {
			// 再试一次 Metadata，确保 Broker 能响应
			_, err = conn.Controller()
			conn.Close()
			if err == nil {
				log.Println("✅ kafka is fully ready")
				return
			}
		}
		log.Printf("⏳ kafka not fully ready, retry %d/40: %v", i, err)
		time.Sleep(2 * time.Second)
	}
	log.Fatal("❌ kafka not ready after retries")
}
