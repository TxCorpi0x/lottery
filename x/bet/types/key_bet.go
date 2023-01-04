package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// ActiveBetKeyPrefix is the active bet prefix to retrieve all active Bet
	ActiveBetKeyPrefix = "Bet/Active/"
	// SettledBetKeyPrefix is the settled bet prefix to retrieve all settled Bet
	SettledBetKeyPrefix = "Bet/Settled/"
)

// ActiveBetKey returns the store key to retrieve an Active Bet from the index fields
func ActiveBetKey(creator string) []byte {
	var key []byte

	key = append(key, []byte(creator)...)
	key = append(key, []byte("/")...)

	return key
}

// SettledBetKey returns the store key to retrieve a Settled Bet from the index fields
func SettledBetKey(lotteryID, betID string) []byte {
	var key []byte

	key = append(key, []byte(lotteryID)...)
	key = append(key, []byte("/")...)

	key = append(key, []byte(betID)...)
	key = append(key, []byte("/")...)

	return key
}
