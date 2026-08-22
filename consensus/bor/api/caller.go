package api

import (
	"context"

	"github.com/photon-storage/bor/common"
	"github.com/photon-storage/bor/common/hexutil"
	"github.com/photon-storage/bor/core/state"
	ethapi "github.com/photon-storage/bor/internal/ethapi"
	"github.com/photon-storage/bor/internal/ethapi/override"
	"github.com/photon-storage/bor/rpc"
)

//go:generate mockgen -destination=./caller_mock.go -package=api . Caller
type Caller interface {
	Call(ctx context.Context, args ethapi.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, overrides *override.StateOverride, blockOverrides *override.BlockOverrides) (hexutil.Bytes, error)
	CallWithState(ctx context.Context, args ethapi.TransactionArgs, blockNrOrHash *rpc.BlockNumberOrHash, state *state.StateDB, overrides *override.StateOverride, blockOverrides *override.BlockOverrides) (hexutil.Bytes, error)
	GetBalance(ctx context.Context, address common.Address, blockNrOrHash rpc.BlockNumberOrHash) (*hexutil.Big, error)
}
