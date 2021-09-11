package services

import (
	"github.com/gomodule/redigo/redis"
	"github.com/tv2169145/redis-ticket-store/localTickit"
	"github.com/tv2169145/redis-ticket-store/remoteTicket"
	"strconv"
)

func BuyTicketByRedis(redisConn redis.Conn, localTicket *localTickit.LocalTicket, redisKeys *remoteTicket.RemoteTicketsKeys) (bool, string) {
	var logMsg string
	if localTicket.LocalDeductionTicket() && redisKeys.RemoteDeductionTicket(redisConn) {
		logMsg = "result:1, 已賣出:" + strconv.FormatInt(localTicket.LocalSales, 10)

		return true, logMsg
	} else {
		logMsg = "result:0, 已賣出:" + strconv.FormatInt(localTicket.LocalSales, 10)

		return false, logMsg
	}
}
