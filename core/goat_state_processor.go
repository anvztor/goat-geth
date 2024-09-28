package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/types/goattypes"
	"github.com/ethereum/go-ethereum/params"
	"github.com/holiman/uint256"
)

var (
	gfBasePoint    = big.NewInt(200)
	gfMaxBasePoint = big.NewInt(1e4)
)

func ProcessGoatFoundationReward(statedb *state.StateDB, gasFees *big.Int) *big.Int {
	if gasFees.BitLen() == 0 {
		return new(big.Int)
	}

	// foundation tax 2%
	tax := new(big.Int).Mul(gasFees, gfBasePoint)
	tax.Div(tax, gfMaxBasePoint)

	if tax.BitLen() != 0 {
		f, _ := uint256.FromBig(tax)
		statedb.AddBalance(goattypes.GoatFoundationContract, f, tracing.BalanceIncreaseRewardTransactionFee)
	}

	// add gas revenue to locking contract
	gas := new(big.Int).Sub(gasFees, tax)
	if gas.BitLen() != 0 {
		f, _ := uint256.FromBig(gas)
		statedb.AddBalance(goattypes.LockingContract, f, tracing.BalanceIncreaseRewardTransactionFee)
	}
	return gas
}

func ProcessGoatRequests(reward *big.Int, logs []*types.Log, config *params.ChainConfig) (types.Requests, error) {
	requests := make(types.Requests, 0, 1)
	requests = append(requests, types.NewRequest(types.NewGoatGasRevenue(reward)))
	for _, log := range logs {
		switch log.Address {
		case goattypes.RelayerContract:
			reqs, err := types.GetRelayerRequests(log.Topics, log.Data)
			if err != nil {
				return nil, err
			}
			requests = append(requests, reqs...)
		case goattypes.BridgeContract:
			reqs, err := types.GetBridgeRequests(log.Topics, log.Data)
			if err != nil {
				return nil, err
			}
			requests = append(requests, reqs...)
		case goattypes.LockingContract:
			reqs, err := types.GetLockingRequests(log.Topics, log.Data)
			if err != nil {
				return nil, err
			}
			requests = append(requests, reqs...)
		}
	}
	return requests, nil
}
