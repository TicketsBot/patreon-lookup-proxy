package config

type (
	Config struct {
		Patreon Patreon
		Server  Server
		Oauth   Oauth
	}

	Patreon struct {
		CampaignId string
		Tiers      []int
	}

	Server struct {
		Host string
		Key  string
	}

	Oauth struct {
		AccessToken string
	}
)
