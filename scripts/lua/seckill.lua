-- seckill_real.lua（最简化版）

math.randomseed(os.time())

request = function()
    -- 直接用随机大整数作为 user_id
    local user_id = math.random(100000000, 999999999)

    return wrk.format(
        "GET",
        "/seckill?user_id=" .. user_id .. "&activity_id=1"
    )
end