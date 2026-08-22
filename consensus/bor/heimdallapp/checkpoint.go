package heimdallapp

import (
	"context"

	"github.com/photon-storage/bor/common"
	"github.com/photon-storage/bor/consensus/bor/heimdall/checkpoint"
	"github.com/photon-storage/bor/log"

	hmTypes "github.com/0xPolygon/heimdall-v2/x/checkpoint/types"
)

func (h *HeimdallAppClient) FetchCheckpointCount(_ context.Context) (int64, error) {
	log.Info("Fetching checkpoint count")

	res, err := h.hApp.CheckpointKeeper.GetAckCount(h.NewContext())
	if err != nil {
		return 0, err
	}

	log.Info("Fetched checkpoint count")

	return int64(res), nil
}

func (h *HeimdallAppClient) FetchCheckpoint(_ context.Context, number int64) (*checkpoint.Checkpoint, error) {
	log.Info("Fetching checkpoint", "number", number)

	res, err := h.hApp.CheckpointKeeper.GetCheckpointByNumber(h.NewContext(), uint64(number))
	if err != nil {
		return nil, err
	}

	log.Info("Fetched checkpoint", "number", number)

	return toBorCheckpoint(res), nil
}

func toBorCheckpoint(hdCheckpoint hmTypes.Checkpoint) *checkpoint.Checkpoint {
	return &checkpoint.Checkpoint{
		Proposer:   common.HexToAddress(hdCheckpoint.Proposer),
		StartBlock: hdCheckpoint.StartBlock,
		EndBlock:   hdCheckpoint.EndBlock,
		RootHash:   common.BytesToHash(hdCheckpoint.RootHash),
		BorChainID: hdCheckpoint.BorChainId,
		Timestamp:  hdCheckpoint.Timestamp,
	}
}
