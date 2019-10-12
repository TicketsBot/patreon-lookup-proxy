package app

import (
	"github.com/mxpv/patreon-go"
	"sync"
)

var (
	Pledges      = make(map[*patreon.User]patreon.Pledge)
	PledgesMutex = sync.Mutex{}
)
