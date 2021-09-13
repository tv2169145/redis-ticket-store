package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tv2169145/redis-ticket-store/api"
	"github.com/tv2169145/redis-ticket-store/middlewares"
)

var (
	Router *gin.Engine
)

func init () {
	Router = gin.New()
	apiGroup := Router.Group("/api").Use(middlewares.RateLimitMiddleware)
	apiGroup.GET("/buy_ticket", api.BuyTicket)
}
