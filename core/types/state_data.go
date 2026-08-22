package types

import "github.com/photon-storage/bor/common"

// StateSyncData represents state received from Ethereum Blockchain
type StateSyncData struct {
	ID       uint64
	Contract common.Address
	Data     []byte
	TxHash   common.Hash // L1 TxHash
}
