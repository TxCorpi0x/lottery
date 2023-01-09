package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// LotteryKeyPrefix is the prefix to retrieve all Lottery
	LotteryKeyPrefix = "Lottery/value/"
)

// LotteryKey returns the store key to retrieve a Lottery from the index fields
func LotteryKey(
	id uint64,
) []byte {
	return uint64ToBytes(id)
}

// uint64ToBytes converts a uint64 into fixed length bytes for use in store keys.
func uint64ToBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(id))
	return bz
}
