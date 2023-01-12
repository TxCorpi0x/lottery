# Lottery

**lottery** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).
Application contains module named `lottery` which is responsible for lottery management and winner detection and payout, and `bet` module that is responsible for receiving bet placement transactions, validation and settlement of the bet messages.

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development. 

## Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com). The default parameters of the `lottery` module is set inside 
the `config.yml`.

## Build and Run manually

In the root of the project issue the followinf command, this will build the binary from the source.

```
make install
```

### Init and Run

Run these commands one after another
```
rm -rf ~/.lottery/
lotteryd init test --chain-id test
```
Open `~/.lottery/config/genesis.json` and change any occurance of `stake` with `utoken` and save the modified content.

```
lotteryd keys add client1
lotteryd add-genesis-account client1 1000000000000utoken
lotteryd gentx client1 10000000utoken --chain-id test --moniker myval
lotteryd collect-gentxs
lotteryd start
```
The chain should start producing block.

## Modules

---

### Lottery Module

This module is designed to handle `end_block` function, the lottery items in the blockchain state are stored with `timestamp` of the block time as key in order to be sort-able and keep a complete history of the lottery objects.

New lottery may add in each `end_block` if no current lottery is present. (active lottery is the latest lottery object that has not any winner yet). 
If the bets of the lottery does not meet these conditions, lottery will continue in the next block:

- Total actve (unsettled) bet count should not be less than 10;
- Operator account of the block proposer should not have an active bet in the state.

#### Wiinning Decision

Winner is chosen randomly according to the hash of active bets and modulo.
```
winnerIndex := (betsHash ^ 0xFFFF) % betCount
```
It gets the hash of bets list ordered ascending according to the bet placement timestamp(Sequencial ID) then gets the first 16 bytes and calculates the modulo of bet items in the list.
the calculated index is used to choose the winner bet.

there is another approach for deciding the winner according to the proposer consensus address to create the hash that replaces `betsHash` in previous formula with `proposerHash` with the sam algorithm.

#### Payout 

Winner should be paid out according to the amount of its bet with following conditions.

- Winner placed highest amount, will receive the entire pool balance (Lottery module account balance)
- Winner placed lowest amount, will receive nothing and all of bet amounts of active bets remoain in the pool balance and lottery will end.
- The rest of conditions, will receive total bet amount of current lottery without fees.

#### Hash Algorithm

This Repo uses [xxhash](http://cyan4973.github.io/xxHash/) algoritm as an efficient algorithm for hashing purposes.

#### Parameters

Lottery module has two parameters that can be changed in the genesis or via governance parameter change proposal(running chain).
`LotteryFee` is being received in each bet placement from the bettor account. and the `BetSize` contains a minimum and maximum allowed amount of the bet placement.
```
lottery_fee: "5000000"
bet_size:
    min_bet: "1000000"
    max_bet: "100000000"
```

---

### Bet Module

#### Placement

This module exposes the `CreateBet` method to get the incomming bets and store it in the active bet list, the active bets are not related to any lottery item until the `lottery` module
end blocker fires and set the winner and the lottery id.

With thid design there is no need to update counter of the lottery and unique bet for the certain bettor account.

The keys are stored with this prefix that enables us to ensure we have only one bet from a user and any bet item that is stored in active bet store is not processed yet.

```
ActiveBetPrefixConstant/{creatorAddress}
```

User has to pay `lottery_fee` and `bet_amount` in each bet placement and if a user tries to bet in a certain lottery multiple times, will have to pay this amount in each transaction but this
will not add new bet to the state and just the current active bet will be modified. This huge amount of loss prevent users to run spam bets on the system.

The sum of `lottery_fee` and `bet_amount` will transfer to `lottery` module account (pool) and subtracted from the bettor balance.

#### Settlement

Settlement will happen in bulk and fired by `lottery` but the logic resides in the bet module. the steps are as follows:

- Lottery decides about the winner
- Removes all of items in the active bet store
- Sets the lottery-related data (lottery id) 
- Creates settled bet objects in the settled bet store of the `bet` module by calling the defined method of the bet `keeper`

The keys are stored with this prefix and enable us to query all of settled bets of finished lottery.

```

SettledBetPrefixConstant/{lotteryID}/{BetID}
```

---

## Demo

There is a `TestDemo` in the `abci_test.go` of the lottery module that simulates lottery for `100` blocks or until all clients run out their tokens.
each client has `500token` uqual to `500000000utoken` and place bet equal to their number. to run the test issue this command:

```
go test -v -timeout 30m -run ^TestDemo$ github.com/vjdmhd/lottery/x/lottery
```

it will print out the current ended block statistics.
