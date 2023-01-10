package types

import (
	"encoding/binary"

	"github.com/vjdmhd/lottery/utils"
)

var _ binary.ByteOrder

var (
	// ActiveBetKeyPrefix is the active bet prefix to retrieve all active Bet
	ActiveBetKeyPrefix = []byte{0x00}
	// SettledBetKeyPrefix is the settled bet prefix to retrieve all settled Bet
	SettledBetKeyPrefix = []byte{0x01}
	// BetStatsKey is the total count of bets key
	BetStatsKey = []byte{0x02}
)

// ActiveBetKey returns the store key to retrieve an Active Bet from the index fields
func ActiveBetKey(creator string) []byte {
	var key []byte

	key = append(key, []byte(creator)...)

	return key
}

// SettledBetKey returns the store key to retrieve a Settled Bet from the index fields
func SettledBetKey(lotteryID, betID uint64) []byte {
	var key []byte

	key = append(key, utils.Uint64ToBytes(lotteryID)...)
	key = append(key, utils.Uint64ToBytes(betID)...)

	return key
}
