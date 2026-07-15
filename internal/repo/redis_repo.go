package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitRedis(addr string) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         addr,
		PoolSize:     20,
		MinIdleConns: 5,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("redis connect failed: %v", err))
	}
}

const seckillLua = `
local stock = tonumber(redis.call("GET", KEYS[1]))
if not stock or stock <= 0 then
    return -1
end
if redis.call("SISMEMBER", KEYS[2], ARGV[1]) == 1 then
    return -2
end
redis.call("DECR", KEYS[1])
redis.call("SADD", KEYS[2], ARGV[1])
return 1
`

var seckillScript = redis.NewScript(seckillLua)

// 返回值：true = 成功, false = 失败, err = 系统异常
func SeckillDecr(ctx context.Context, userID, activityID string) (bool, error) {
	stockKey := "seckill:stock:" + activityID
	userKey := "seckill:users:" + activityID

	res, err := seckillScript.Run(ctx, rdb, []string{stockKey, userKey}, userID).Int64()
	if err != nil {
		return false, err
	}

	switch res {
	case 1:
		return true, nil
	case -1:
		return false, nil // 库存不足
	case -2:
		return false, nil // 重复抢单
	default:
		return false, fmt.Errorf("unexpected lua result: %d", res)
	}
}
