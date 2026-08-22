package wit

import (
	"github.com/photon-storage/bor/core"
	"github.com/photon-storage/bor/core/forkid"
	"github.com/photon-storage/bor/rlp"
)

// enrEntry is the ENR entry which advertises `wit` protocol on the discovery.
type enrEntry struct {
	ForkID forkid.ID // Fork identifier per EIP-2124

	// Ignore additional fields (for forward compatibility).
	Rest []rlp.RawValue `rlp:"tail"`
}

// ENRKey implements enr.Entry.
func (e enrEntry) ENRKey() string {
	return "wit"
}

// currentENREntry constructs an `eth` ENR entry based on the current state of the chain.
func currentENREntry(chain *core.BlockChain) *enrEntry {
	head := chain.CurrentHeader()

	return &enrEntry{
		ForkID: forkid.NewID(chain.Config(), chain.Genesis(), head.Number.Uint64(), head.Time),
	}
}
