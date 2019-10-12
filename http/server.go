package http

import (
	"github.com/TicketsBot/patreon-lookup-proxy/config"
	"github.com/TicketsBot/patreon-lookup-proxy/http/endpoints"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	router.GET("/ispremium", endpoints.IsPremium)

	if err := router.Run(config.Conf.Server.Host); err != nil {
		panic(err)
	}
}
