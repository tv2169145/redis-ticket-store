package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/tv2169145/redis-ticket-store/localTickit"
	"github.com/tv2169145/redis-ticket-store/redisConnection"
	redisTicket "github.com/tv2169145/redis-ticket-store/remoteTicket"
	"github.com/tv2169145/redis-ticket-store/services"
	"github.com/tv2169145/redis-ticket-store/utils"
	"net/http"
)

var (
	redisPool *redis.Pool
	done chan int
	LocalTicket *localTickit.LocalTicket
	RemoteTicket *redisTicket.RemoteTicketsKeys
)

func init() {
	LocalTicket = localTickit.NewLocalTicket(100, 0)
	RemoteTicket = redisTicket.NewRemoteTicketKeys()
	redisPool = redisConnection.NewPool()
	done = make(chan int, 1)
	done<-1
}

func BuyTicket(c *gin.Context) {
	redisConn := redisPool.Get()
	//全局讀寫鎖
	<-done
	result, msg := services.BuyTicketByRedis(redisConn, LocalTicket, RemoteTicket)
	done <- 1
	//將狀態寫入log中
	utils.WriteLog(msg, "./logs/stat.log")
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"msg": msg,
		"result": result,
	})
}
