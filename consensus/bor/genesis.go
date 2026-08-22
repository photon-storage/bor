package bor

import (
	"math/big"

	"github.com/photon-storage/bor/common"
	"github.com/photon-storage/bor/consensus/bor/clerk"
	"github.com/photon-storage/bor/consensus/bor/statefull"
	"github.com/photon-storage/bor/core/state"
	"github.com/photon-storage/bor/core/types"
	"github.com/photon-storage/bor/core/vm"
)

//go:generate mockgen -destination=./genesis_contract_mock.go -package=bor . GenesisContract
type GenesisContract interface {
	CommitState(event *clerk.EventRecordWithTime, state vm.StateDB, header *types.Header, chCtx statefull.ChainContext, vmCfg vm.Config) (uint64, error)
	LastStateId(state *state.StateDB, number uint64, hash common.Hash) (*big.Int, error)
}
