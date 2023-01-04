package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// BetKeyPrefix is the prefix to retrieve all Bet
	BetKeyPrefix = "Bet/value/"
)

// BetKey returns the store key to retrieve a Bet from the index fields
func BetKey(
	lotteryID, creator string,
) []byte {
	var key []byte

	lotteryIDBytes := []byte(lotteryID)
	key = append(key, lotteryIDBytes...)
	key = append(key, []byte("/")...)

	creatorBytes := []byte(creator)
	key = append(key, creatorBytes...)
	key = append(key, []byte("/")...)

	return key
}
