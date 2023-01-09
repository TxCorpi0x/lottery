package params

// App Parameters
const (
	// HumanCoinUnit is human readable representation of the coin name
	HumanCoinUnit = "token"
	// BaseCoinUnit is the actual name of coin used in transaction
	BaseCoinUnit = "utoken"
	// SGEExponent is the exponential digits of the coin
	LOTExponent = 6
	// DefaultBondDenom is the default staking denom of application
	DefaultBondDenom = BaseCoinUnit
)
