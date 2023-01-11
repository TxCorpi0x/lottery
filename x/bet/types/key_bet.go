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
	return KeyPrefix(creator)
}

// SettledBetOfLotteryPrefix returns the store key to retrieve all Settled Bet of lottery
func SettledBetOfLotteryPrefix(lotteryID uint64) []byte {
	return append(SettledBetKeyPrefix, utils.Uint64ToBytes(lotteryID)...)
}

// SettledBetKey returns the store key to retrieve a Settled Bet from the index fields
func SettledBetKey(lotteryID, betID uint64) []byte {
	return append(utils.Uint64ToBytes(lotteryID), utils.Uint64ToBytes(betID)...)
}
