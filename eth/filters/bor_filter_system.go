package filters

import (
	"time"

	"github.com/photon-storage/bor/core"
	"github.com/photon-storage/bor/core/types"
	"github.com/photon-storage/bor/rpc"
)

func (es *EventSystem) handleStateSyncEvent(filters filterIndex, ev core.StateSyncEvent) {
	for _, f := range filters[StateSyncSubscription] {
		f.stateSyncData <- ev.Data
	}
}

// SubscribeNewDeposits creates a subscription that writes details about the new state sync events (from mainchain to Bor)
func (es *EventSystem) SubscribeNewDeposits(data chan *types.StateSyncData) *Subscription {
	sub := &subscription{
		id:            rpc.NewID(),
		typ:           StateSyncSubscription,
		created:       time.Now(),
		logs:          make(chan []*types.Log),
		txs:           make(chan []*types.Transaction),
		headers:       make(chan *types.Header),
		stateSyncData: data,
		installed:     make(chan struct{}),
		err:           make(chan error),
	}

	return es.subscribe(sub)
}
