package bor

import (
	"context"

	borTypes "github.com/0xPolygon/heimdall-v2/x/bor/types"
	stakeTypes "github.com/0xPolygon/heimdall-v2/x/stake/types"

	"github.com/photon-storage/bor/common"
	"github.com/photon-storage/bor/consensus/bor/valset"
	"github.com/photon-storage/bor/core"
	"github.com/photon-storage/bor/core/state"
	"github.com/photon-storage/bor/core/types"
	"github.com/photon-storage/bor/core/vm"
	"github.com/photon-storage/bor/rpc"
)

//go:generate mockgen -destination=./span_mock.go -package=bor . Spanner
type Spanner interface {
	GetCurrentSpan(ctx context.Context, headerHash common.Hash, state *state.StateDB) (*borTypes.Span, error)
	GetCurrentValidatorsByHash(ctx context.Context, headerHash common.Hash, blockNumber uint64) ([]*valset.Validator, error)
	GetCurrentValidatorsByBlockNrOrHash(ctx context.Context, blockNrOrHash rpc.BlockNumberOrHash, blockNumber uint64) ([]*valset.Validator, error)
	CommitSpan(ctx context.Context, minimalSpan borTypes.Span, validators, producers []stakeTypes.MinimalVal, state vm.StateDB, header *types.Header, chainContext core.ChainContext, vmCfg vm.Config) error
}
