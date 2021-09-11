package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tv2169145/redis-ticket-store/api"
)

var (
	Router *gin.Engine
)

func init () {
	Router = gin.New()
	Router.GET("/buy_ticket", api.BuyTicket)
}
