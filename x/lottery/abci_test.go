package lottery_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/require"
	"github.com/vjdmhd/lottery/app/params"
	lotsim "github.com/vjdmhd/lottery/testutil/simapp"
	betKeeper "github.com/vjdmhd/lottery/x/bet/keeper"
	bettypes "github.com/vjdmhd/lottery/x/bet/types"
	lotteryModule "github.com/vjdmhd/lottery/x/lottery"
	"github.com/vjdmhd/lottery/x/lottery/types"
)

func TestDemo(t *testing.T) {

	tApp, ctx, err := lotsim.GetTestObjects()
	if err != nil {
		panic(err)
	}
	betMsgSrv := betKeeper.NewMsgServerImpl(tApp.BetKeeper)

	clientInices := []int{}
	validators := []lotsim.TestValidator{lotsim.TestParamValidatorAddresses["val1"], lotsim.TestParamValidatorAddresses["val2"]}

	for blockNo := 1; blockNo <= 100; blockNo++ {
		fmt.Printf("---------------------------Block %d Ended------------------------------\n", blockNo)

		// pick random validator between two validators (val1, val2)
		rand.Seed(time.Now().UnixNano())
		proposer := validators[rand.Intn(len(validators))]

		// set the current proposer
		ctx = ctx.
			WithProposer(proposer.ConsAddress).
			WithBlockTime(time.Now()).
			WithBlockHeight(int64(blockNo))

		// refill the indices slice
		for i := 2; i <= 21; i++ {
			cl := lotsim.TestParamUsers["client"+cast.ToString(i)]
			balance := tApp.BankKeeper.SpendableCoins(ctx, cl.Address).AmountOf(params.DefaultBondDenom)

			// check if the balance is sufficient to be available in the potential client list
			betAmountAndFee := int64(i * cast.ToInt(math.Pow(10, params.LOTExponent)))
			totalAmount := tApp.LotteryKeeper.GetLotteryFee(ctx).Add(sdk.NewInt(betAmountAndFee))
			totalCoin := sdk.NewCoin(params.DefaultBondDenom, totalAmount)
			if balance.GT(totalCoin.Amount) {
				clientInices = append(clientInices, i)
			}
		}

		// there is not enough bets to continue blocks
		// note: if this gets commented, the blocks loop will continue till 100
		// 		 and clients will pay lottery_fee and bet_amount until the account balance gets lower than sum of fee and bet amount
		if len(clientInices) < types.MinBetCount {
			break
		}

		// remove random element from inices slice until running out
		for len(clientInices) > 0 {
			// pick a random index
			rand.Seed(time.Now().UnixNano())
			reandomIndex := rand.Intn(len(clientInices))

			// number of the client is randomly picked and
			clientNumber := clientInices[reandomIndex]

			// remove picked index from slice
			clientInices = removeFromIntSlice(clientInices, reandomIndex)

			// // skip the first client because the client1 and client 2
			// // are validator operators, we skip the first one to decrease possibility of
			// // being proposer
			// clientNumber += 1

			// get creator address from client
			creator := lotsim.TestParamUsers["client"+cast.ToString(clientNumber)]

			// create bet by running message server method
			betAmount := clientNumber * cast.ToInt(math.Pow(10, params.LOTExponent))
			// fmt.Println(creator.Address.String())
			_, err := betMsgSrv.CreateBet(ctx, bettypes.NewMsgCreateBet(
				creator.Address.String(),
				sdk.NewInt(int64(betAmount)),
			))

			require.NoError(t, err)

			// sleep to allow random number generating
			time.Sleep(50 * time.Microsecond)
		}

		// wait and simulate the block time
		time.Sleep(5 * time.Second)

		// run end blocker to simulate end block
		lotteryModule.EndBlocker(ctx, tApp.LotteryKeeper)

	}

	finalActiveBets := tApp.BetKeeper.GetAllActiveBet(ctx)
	require.LessOrEqual(t, len(finalActiveBets), types.MinBetCount)

}

// removes element and returns the slice
func removeFromIntSlice(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
