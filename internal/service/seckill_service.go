package service

import (
	"context"
	"seckill/internal/repo"
)

func DoSeckill(ctx context.Context, userID, activityID string) (int64, error) {
	// 1. Redis 原子扣减（Lua）
	success, err := repo.SeckillDecr(ctx, userID, activityID)
	if err != nil || !success {
		return -1, err
	}

	// 2. 发送 Kafka（异步落单）
	orderID, err := repo.SendOrderEvent(ctx, userID, activityID)
	if err != nil {
		return -1, err
	}

	return orderID, nil
}
