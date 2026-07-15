package service

import (
	"context"
	"errors"
	"seckill/internal/repo"
)

var (
	ErrSoldOut = errors.New("sold out")
)

func DoSeckill(ctx context.Context, userID, activityID string) error {
	success, err := repo.SeckillDecr(ctx, userID, activityID)
	if err != nil {
		return err
	}
	if !success {
		return ErrSoldOut
	}
	return repo.SendOrderEvent(ctx, userID, activityID)
}
