package types

import (
	"encoding/binary"

	"github.com/vjdmhd/lottery/utils"
)

var _ binary.ByteOrder

var (
	// LotteryKeyPrefix is the prefix to retrieve all Lottery
	LotteryKeyPrefix = []byte{0x00}
)

// LotteryKey returns the store key to retrieve a Lottery from the index fields
func LotteryKey(id uint64) []byte {
	return utils.Uint64ToBytes(id)
}
