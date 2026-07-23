package service

import (
	"context"
	"errors"
	repo "seckill/internal/repo"

	"github.com/google/uuid"
)

var ErrSoldOut = errors.New("sold out")

func DoSeckill(ctx context.Context, userID, activityID string) (string, error) {
	orderID := uuid.NewString()
	res, err := repo.Seckill(ctx, userID, activityID, orderID)
	if err != nil {
		return "", err
	}
	switch res {
	case repo.SeckillOK:
		return orderID, nil
	case repo.SeckillSoldOut, repo.SeckillRepeated:
		return "", ErrSoldOut
	default:
		return "", errors.New("unknown error")
	}
}
