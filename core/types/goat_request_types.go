package types

import "github.com/ethereum/go-ethereum/common"

const (
	GoatGasRevenueRequestType = iota + 0x60
	GoatAddVoterRequestType
	GoatRemoveVoterRequestType
	GoatWithdrawalRequestType
	GoatReplaceByFeeRequestType
	GoatCancel1RequestType
)

var (
	GoatAddVoterTopoic    = common.HexToHash("0x101c617f43dd1b8a54a9d747d9121bbc55e93b88bc50560d782a79c4e28fc838")
	GoatRemoveVoterTopic  = common.HexToHash("0x183393fc5cffbfc7d03d623966b85f76b9430f42d3aada2ac3f3deabc78899e8")
	GoatWithdrawalTopic   = common.HexToHash("0xbe7c38d37e8132b1d2b29509df9bf58cf1126edf2563c00db0ef3a271fb9f35b")
	GoatReplaceByFeeTopic = common.HexToHash("0x19875a7124af51c604454b74336ce2168c45bceade9d9a1e6dfae9ba7d31b7fa")
	GoatCancel1Topic      = common.HexToHash("0x0106f4416537efff55311ef5e2f9c2a48204fcf84731f2b9d5091d23fc52160c")
)
