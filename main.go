package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tv2169145/redis-ticket-store/routers"
)

var (
	router *gin.Engine
)

func init() {
	router = routers.Router
}

func main() {
	router.Run(":8888")
}
