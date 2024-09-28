package types

import (
	"encoding/json"
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

const (
	GoatGasRevenueRequestType byte = iota + 0x60
	GoatAddVoterRequestType
	GoatRemoveVoterRequestType
	GoatWithdrawalRequestType
	GoatReplaceByFeeRequestType
	GoatCancel1RequestType
	GoatCreateValidatorType
	GoatLockType
	GoatUnlockType
	GoatClaimType
	GoatSetTokenWeight
)

var (
	GoatAddVoterTopoic    = common.HexToHash("0x101c617f43dd1b8a54a9d747d9121bbc55e93b88bc50560d782a79c4e28fc838")
	GoatRemoveVoterTopic  = common.HexToHash("0x183393fc5cffbfc7d03d623966b85f76b9430f42d3aada2ac3f3deabc78899e8")
	GoatWithdrawalTopic   = common.HexToHash("0xbe7c38d37e8132b1d2b29509df9bf58cf1126edf2563c00db0ef3a271fb9f35b")
	GoatReplaceByFeeTopic = common.HexToHash("0x19875a7124af51c604454b74336ce2168c45bceade9d9a1e6dfae9ba7d31b7fa")
	GoatCancel1Topic      = common.HexToHash("0x0106f4416537efff55311ef5e2f9c2a48204fcf84731f2b9d5091d23fc52160c")

	GoatCreateValidatorTopic   = common.HexToHash("0xf3aa84440b70359721372633122645674adb6dbb72622a222627248ef053a7dd")
	GoatValidatorLockTopic     = common.HexToHash("0xec36c0364d931187a76cf66d7eee08fad0ec2e8b7458a8d8b26b36769d4d13f3")
	GoatValidatorUnlockTopic   = common.HexToHash("0x40f2a8c5e2e2a9ad2f4e4dfc69825595b526178445c3eb22b02edfd190601db7")
	GoatValidatorClaimTopic    = common.HexToHash("0xa983a6cfc4bd1095dac7b145ae020ba08e16cc7efa2051cc6b77e4011b9ee99b")
	GoatUpdateTokenWeightTopic = common.HexToHash("0xb59bf4596e5415117fb4625044cb5b0ca5b273742825b026d06afe82a48e6217")
)

func (r *Request) ForGoat() bool {
	switch r.inner.requestType() {
	case DepositRequestType:
		return false
	default:
		return true
	}
}

// Request intermediate type for json codec
type requestMarshaling struct {
	Type byte            `json:"type"`
	Data json.RawMessage `json:"data"`
}

// UnmarshalJSON implements json.Marshaler interface
func (r *Request) MarshalJSON() ([]byte, error) {
	if r.inner == nil {
		return nil, errors.New("no request data")
	}

	data, err := json.Marshal(r.inner)
	if err != nil {
		return nil, err
	}

	return json.Marshal(requestMarshaling{Type: r.inner.requestType(), Data: data})
}

// UnmarshalJSON implements json.Unmarshaler interface
func (r *Request) UnmarshalJSON(b []byte) error {
	var raw requestMarshaling
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case GoatGasRevenueRequestType:
		r.inner = new(GasRevenue)
	case GoatAddVoterRequestType:
		r.inner = new(AddVoter)
	case GoatRemoveVoterRequestType:
		r.inner = new(RemoveVoter)
	case GoatWithdrawalRequestType:
		r.inner = new(BridgeWithdrawal)
	case GoatReplaceByFeeRequestType:
		r.inner = new(ReplaceByFee)
	case GoatCancel1RequestType:
		r.inner = new(Cancel1)
	case GoatCreateValidatorType:
		r.inner = new(CreateValidator)
	case GoatLockType:
		r.inner = new(ValidatorLock)
	case GoatUnlockType:
		r.inner = new(ValidatorUnlock)
	case GoatClaimType:
		r.inner = new(GoatRewardClaim)
	case GoatSetTokenWeight:
		r.inner = new(SetTokenWeight)

	case DepositRequestType:
		r.inner = new(Deposit)
	default:
		return ErrRequestTypeNotSupported
	}

	return json.Unmarshal(raw.Data, r.inner)
}
