package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/tv2169145/redis-ticket-store/redisConnection"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	KEY_PREFIX = "client_ip"
	LUA_SCRIPT = `
	local times = redis.call('incr', KEYS[1])
	if times == 1 then
	redis.call('expire', KEYS[1], ARGV[1]) end
	if times > tonumber(ARGV[2]) then 
	return 0
	end
	return times
`
)

var (
	limit = 10
	expired = 60  //單位(秒)->一分鐘10次
	timeZone, _ = time.LoadLocation("Asia/Taipei")
)

func RateLimitMiddleware(c *gin.Context) {
	log.Println("middleware here...")
	clientIp := c.ClientIP()
	redisKey := KEY_PREFIX + clientIp
	log.Println(clientIp, redisKey)

	// 檢查該ip key 是否存在於redis
	redisConn := redisConnection.GetRedisConn()
	luaScript := redis.NewScript(1, LUA_SCRIPT)
	times, err := redis.Int(luaScript.Do(redisConn, redisKey, expired, limit))
	if err != nil {
		log.Println("LUA Error", err)
	}


	if times != 0 {
		//resetSeconds, _ := redis.Int(redisConn.Do("TTL", clientIp))
		//resetAt := time.Now().In(timeZone).Add(time.Second*time.Duration(resetSeconds))
		//c.Header("X-RateLimit-Reset", resetAt.String())
		c.Header("X-RateLimit-Max", strconv.Itoa(limit))
		c.Header("X-RateLimit-Remaining", strconv.Itoa(limit-times))
		c.Next()
	} else {
		c.JSON(http.StatusTooManyRequests, nil)
	}
}