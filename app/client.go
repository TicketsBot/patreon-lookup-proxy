package app

import (
	"context"
	"github.com/TicketsBot/patreon-go"
	"github.com/TicketsBot/patreon-lookup-proxy/config"
	"golang.org/x/oauth2"
)

func NewPatreonClient() *patreon.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: config.Conf.Oauth.AccessToken})
	tc := oauth2.NewClient(context.Background(), ts)
	
	client := patreon.NewClient(tc)

	return client
}
