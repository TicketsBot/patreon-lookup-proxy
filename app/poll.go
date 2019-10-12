package app

import (
	"github.com/mxpv/patreon-go"
	"github.com/TicketsBot/patreon-lookup-proxy/config"
	"github.com/apex/log"
	"time"
)

func Poll(client *patreon.Client) {
	cursor := ""

	pledges := make(map[*patreon.User]patreon.Pledge)

	for {
		res, err := client.FetchPledges(config.Conf.Patreon.CampaignId, patreon.WithPageSize(25), patreon.WithCursor(cursor))
		if err != nil {
			time.Sleep(time.Second)
			log.Error(err.Error())
			continue
		}

		users := make(map[string]*patreon.User)
		for _, item := range res.Included.Items {
			user, ok := item.(*patreon.User)
			if !ok {
				continue
			}

			users[user.ID] = user
		}

		for _, pledge := range res.Data {
			user, ok := users[pledge.Relationships.Patron.Data.ID]
			if !ok {
				continue
			}

			pledges[user] = pledge
		}

		next := res.Links.Next
		if next == "" {
			break
		}

		cursor = next
	}

	PledgesMutex.Lock()
	Pledges = pledges
	PledgesMutex.Unlock()
}
