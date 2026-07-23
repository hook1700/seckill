package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitRedis(addr string, poolSize int) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         addr,
		PoolSize:     poolSize,
		MinIdleConns: poolSize / 4,
		DialTimeout:  2 * time.Second,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("redis connect failed: %v", err))
	}
}

var seckillScript = redis.NewScript(`
local stock = tonumber(redis.call("GET", KEYS[1]))
if not stock or stock <= 0 then
    return -1
end
if redis.call("SISMEMBER", KEYS[2], ARGV[1]) == 1 then
    return -2
end
redis.call("DECR", KEYS[1])
redis.call("SADD", KEYS[2], ARGV[1])
redis.call("LPUSH", KEYS[3], ARGV[1] .. ":" .. ARGV[2])
return 1
`)

func Seckill(ctx context.Context, userID, activityID, orderID string) (int64, error) {
	stockKey := "seckill:stock:" + activityID
	userKey := "seckill:users:" + activityID
	queueKey := "seckill:orders:" + activityID

	res, err := seckillScript.Run(
		ctx,
		rdb,
		[]string{stockKey, userKey, queueKey},
		userID, orderID,
	).Int64()

	if err != nil {
		return 0, err
	}
	return res, nil
}

func PopOrder(ctx context.Context, activityID string) (userID, orderID string, ok bool) {
	res, err := rdb.BRPop(ctx, 5*time.Second, "seckill:orders:"+activityID).Result()
	if err != nil || len(res) < 2 {
		return "", "", false
	}
	parts := splitOrderMsg(res[1])
	if len(parts) != 2 {
		return "", "", false
	}
	return parts[0], parts[1], true
}

func splitOrderMsg(msg string) []string {
	for i := len(msg) - 1; i >= 0; i-- {
		if msg[i] == ':' {
			return []string{msg[:i], msg[i+1:]}
		}
	}
	return nil
}
