// Code generated by "stringer -type=BalanceChangeReason -output gen_balance_change_reason_stringer.go"; DO NOT EDIT.

package tracing

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[BalanceChangeUnspecified-0]
	_ = x[BalanceIncreaseRewardMineUncle-1]
	_ = x[BalanceIncreaseRewardMineBlock-2]
	_ = x[BalanceIncreaseWithdrawal-3]
	_ = x[BalanceIncreaseGenesisBalance-4]
	_ = x[BalanceIncreaseRewardTransactionFee-5]
	_ = x[BalanceDecreaseGasBuy-6]
	_ = x[BalanceIncreaseGasReturn-7]
	_ = x[BalanceIncreaseDaoContract-8]
	_ = x[BalanceDecreaseDaoAccount-9]
	_ = x[BalanceChangeTransfer-10]
	_ = x[BalanceChangeTouchAccount-11]
	_ = x[BalanceIncreaseSelfdestruct-12]
	_ = x[BalanceDecreaseSelfdestruct-13]
	_ = x[BalanceDecreaseSelfdestructBurn-14]
	_ = x[BalanceGoatDepoist-200]
	_ = x[BalanceGoatTax-201]
}

const (
	_BalanceChangeReason_name_0 = "BalanceChangeUnspecifiedBalanceIncreaseRewardMineUncleBalanceIncreaseRewardMineBlockBalanceIncreaseWithdrawalBalanceIncreaseGenesisBalanceBalanceIncreaseRewardTransactionFeeBalanceDecreaseGasBuyBalanceIncreaseGasReturnBalanceIncreaseDaoContractBalanceDecreaseDaoAccountBalanceChangeTransferBalanceChangeTouchAccountBalanceIncreaseSelfdestructBalanceDecreaseSelfdestructBalanceDecreaseSelfdestructBurn"
	_BalanceChangeReason_name_1 = "BalanceGoatDepoistBalanceGoatTax"
)

var (
	_BalanceChangeReason_index_0 = [...]uint16{0, 24, 54, 84, 109, 138, 173, 194, 218, 244, 269, 290, 315, 342, 369, 400}
	_BalanceChangeReason_index_1 = [...]uint8{0, 18, 32}
)

func (i BalanceChangeReason) String() string {
	switch {
	case i <= 14:
		return _BalanceChangeReason_name_0[_BalanceChangeReason_index_0[i]:_BalanceChangeReason_index_0[i+1]]
	case 200 <= i && i <= 201:
		i -= 200
		return _BalanceChangeReason_name_1[_BalanceChangeReason_index_1[i]:_BalanceChangeReason_index_1[i+1]]
	default:
		return "BalanceChangeReason(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
