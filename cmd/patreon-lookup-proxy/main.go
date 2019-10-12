package main

import (
	"github.com/TicketsBot/patreon-lookup-proxy/app"
	"github.com/TicketsBot/patreon-lookup-proxy/config"
	"github.com/TicketsBot/patreon-lookup-proxy/http"
	"time"
)

func main() {
	config.LoadConfig()

	client := app.NewPatreonClient()

	go func() {
		for {
			app.Poll(client)
			time.Sleep(time.Minute)
		}
	}()

	http.StartServer()
}
