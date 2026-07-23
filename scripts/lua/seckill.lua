-- KEYS[1]: stockKey
-- KEYS[2]: userSetKey
-- KEYS[3]: orderQueueKey
-- ARGV[1]: userId
-- ARGV[2]: orderId

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