package service

import (
	"context"
	"errors"
	"seckill/internal/repo"
)

var (
	ErrSoldOut  = errors.New("sold out")
	ErrRepeated = errors.New("already seckilled")
)

func DoSeckill(ctx context.Context, userID, activityID string) error {
	success, err := repo.SeckillDecr(ctx, userID, activityID)
	if err != nil {
		return err // Redis 异常
	}
	if !success {
		return ErrSoldOut // 库存不足或重复
	}
	return nil
}
