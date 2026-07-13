package repo

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var seckillScript *redis.Script

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

func InitRedis(addr string) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		PoolSize: 100,
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(err)
	}

	seckillScript = redis.NewScript(seckillLua)
}
