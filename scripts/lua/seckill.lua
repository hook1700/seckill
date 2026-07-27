-- KEYS[1]: stockKey
-- KEYS[2]: userSetKey
-- KEYS[3]: orderQueueKey
-- ARGV[1]: userId
-- ARGV[2]: orderId

counter = 0
thread_id = 0

setup = function(thread)
    thread_id = thread_id + 1
end

request = function()
    counter = counter + 1
    local user_id = string.format(
        "%d_%d_%d",
        thread_id,
        os.clock() * 1000000,
        counter
    )
    return wrk.format("GET", "/seckill?user_id=" .. user_id .. "&activity_id=1")
end