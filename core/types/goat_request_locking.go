package types

import (
	"bytes"
	"fmt"
	"math/big"
	"slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

//go:generate go run github.com/fjl/gencodec -type GasRevenue -field-override gasRevenueMarshaling -out gen_goat_request_gas_revenue.go
type GasRevenue struct {
	Amount *big.Int `json:"amount"`
}

type gasRevenueMarshaling struct {
	Amount *hexutil.Big
}

type GasRevenues []*GasRevenue

func (s GasRevenues) Len() int { return len(s) }

func (s GasRevenues) EncodeIndex(i int, w *bytes.Buffer) {
	rlp.Encode(w, s[i])
}

// Requests creates a deep copy of each deposit and returns a slice of the
// GasRevenues requests as Request objects.
func (s GasRevenues) Requests() (reqs Requests) {
	for _, d := range s {
		reqs = append(reqs, NewRequest(d))
	}
	return
}

func NewGoatGasRevenue(amount *big.Int) *GasRevenue {
	return &GasRevenue{Amount: amount}
}

func (d *GasRevenue) requestType() byte            { return GoatGasRevenueRequestType }
func (d *GasRevenue) encode(b *bytes.Buffer) error { return rlp.Encode(b, d) }
func (d *GasRevenue) decode(input []byte) error    { return rlp.DecodeBytes(input, d) }
func (d *GasRevenue) copy() RequestData {
	return &GasRevenue{
		Amount: new(big.Int).Set(d.Amount),
	}
}

//go:generate go run github.com/fjl/gencodec -type CreateValidator -field-override createValidatorMarshaling -out gen_goat_request_create_validator.go
type CreateValidator struct {
	Pubkey []byte `json:"pubkey"`
}

type createValidatorMarshaling struct {
	Pubkey hexutil.Bytes
}

func UnpackIntoCreateValidator(data []byte) (*CreateValidator, error) {
	if len(data) != 128 {
		return nil, fmt.Errorf("CreateValidator wrong length: want 128, have %d", len(data))
	}
	return &CreateValidator{Pubkey: data[64:]}, nil
}

func (d *CreateValidator) requestType() byte            { return GoatCreateValidatorType }
func (d *CreateValidator) encode(b *bytes.Buffer) error { return rlp.Encode(b, d) }
func (d *CreateValidator) decode(input []byte) error    { return rlp.DecodeBytes(input, d) }
func (d *CreateValidator) copy() RequestData {
	return &CreateValidator{
		Pubkey: slices.Clone(d.Pubkey),
	}
}

//go:generate go run github.com/fjl/gencodec -type GoatLock -field-override goatLockMarshaling -out gen_goat_request_goat_lock.go
type GoatLock struct {
	Validator common.Address `json:"validator"`
	Token     common.Address `json:"token"`
	Amount    *big.Int       `json:"amount"`
}

type goatLockMarshaling struct {
	Amount *hexutil.Big
}

func (d *GoatLock) requestType() byte            { return GoatLockRequestType }
func (d *GoatLock) encode(b *bytes.Buffer) error { return rlp.Encode(b, d) }
func (d *GoatLock) decode(input []byte) error    { return rlp.DecodeBytes(input, d) }
func (d *GoatLock) copy() RequestData {
	return &GoatLock{
		Validator: d.Validator,
		Token:     d.Token,
		Amount:    new(big.Int).Set(d.Amount),
	}
}

func UnpackIntoValidatorLock(data []byte) (*GoatLock, error) {
	if len(data) != 96 {
		return nil, fmt.Errorf("ValidatorLock wrong length: want 96, have %d", len(data))
	}
	return &GoatLock{
		Validator: common.BytesToAddress(data[:32]),
		Token:     common.BytesToAddress(data[32:64]),
		Amount:    new(big.Int).SetBytes(data[64:]),
	}, nil
}

//go:generate go run github.com/fjl/gencodec -type GoatUnlock -field-override goatUnlockMarshaling -out gen_goat_request_validator_unlock.go
type GoatUnlock struct {
	Id        uint64         `json:"id"`
	Validator common.Address `json:"validator"`
	Token     common.Address `json:"token"`
	Recipient common.Address `json:"recipient"`
	Amount    *big.Int       `json:"amount"`
}

type goatUnlockMarshaling struct {
	Id     hexutil.Uint64
	Amount *hexutil.Big
}

func (d *GoatUnlock) requestType() byte            { return GoatUnlockRequestType }
func (d *GoatUnlock) encode(b *bytes.Buffer) error { return rlp.Encode(b, d) }
func (d *GoatUnlock) decode(input []byte) error    { return rlp.DecodeBytes(input, d) }
func (d *GoatUnlock) copy() RequestData {
	return &GoatUnlock{
		Id:        d.Id,
		Validator: d.Validator,
		Token:     d.Token,
		Recipient: d.Recipient,
		Amount:    new(big.Int).Set(d.Amount),
	}
}

func UnpackIntoValidatorUnlock(data []byte) (*GoatUnlock, error) {
	if len(data) != 160 {
		return nil, fmt.Errorf("ValidatorUnlock wrong length: want 160, have %d", len(data))
	}
	return &GoatUnlock{
		Id:        new(big.Int).SetBytes(data[:32]).Uint64(),
		Validator: common.BytesToAddress(data[32:64]),
		Recipient: common.BytesToAddress(data[64:96]),
		Token:     common.BytesToAddress(data[96:128]),
		Amount:    new(big.Int).SetBytes(data[128:160]),
	}, nil
}

//go:generate go run github.com/fjl/gencodec -type GoatClaimReward -field-override goatRewardClaimMarshaling -out gen_goat_request_reward_claim.go
type GoatClaimReward struct {
	Id        uint64         `json:"id"`
	Validator common.Address `json:"validator"`
	Recipient common.Address `json:"recipient"`
}

type goatRewardClaimMarshaling struct {
	Id hexutil.Uint64
}

func (d *GoatClaimReward) requestType() byte            { return GoatClaimRewardRequestType }
func (d *GoatClaimReward) encode(b *bytes.Buffer) error { return rlp.Encode(b, d) }
func (d *GoatClaimReward) decode(input []byte) error    { return rlp.DecodeBytes(input, d) }
func (d *GoatClaimReward) copy() RequestData {
	return &GoatClaimReward{
		Id:        d.Id,
		Validator: d.Validator,
		Recipient: d.Recipient,
	}
}

func UnpackIntoGoatRewardClaim(data []byte) (*GoatClaimReward, error) {
	if len(data) != 96 {
		return nil, fmt.Errorf("GoatRewardClaim wrong length: want 96, have %d", len(data))
	}
	return &GoatClaimReward{
		Id:        new(big.Int).SetBytes(data[:32]).Uint64(),
		Validator: common.BytesToAddress(data[32:64]),
		Recipient: common.BytesToAddress(data[64:96]),
	}, nil
}

type UpdateTokenWeight struct {
	Token  common.Address `json:"token"`
	Weight uint64         `json:"weight"`
}

func UnpackIntoUpdateTokenWeight(data []byte) (*UpdateTokenWeight, error) {
	if len(data) != 64 {
		return nil, fmt.Errorf("UpdateTokenWeight wrong length: want 64, have %d", len(data))
	}
	return &UpdateTokenWeight{
		Token:  common.BytesToAddress(data[:32]),
		Weight: new(big.Int).SetBytes(data[32:64]).Uint64(),
	}, nil
}

func (d *UpdateTokenWeight) requestType() byte            { return GoatUpdateTokenWeightRequestType }
func (d *UpdateTokenWeight) encode(b *bytes.Buffer) error { return rlp.Encode(b, d) }
func (d *UpdateTokenWeight) decode(input []byte) error    { return rlp.DecodeBytes(input, d) }
func (d *UpdateTokenWeight) copy() RequestData {
	return &UpdateTokenWeight{
		Token:  d.Token,
		Weight: d.Weight,
	}
}

//go:generate go run github.com/fjl/gencodec -type UpdateTokenThreshold -field-override updateTokenThresholdMarshaling -out gen_goat_request_update_token_threshold.go
type UpdateTokenThreshold struct {
	Token     common.Address `json:"token"`
	Threshold *big.Int       `json:"threshold"`
}

type updateTokenThresholdMarshaling struct {
	Threshold *hexutil.Big
}

func UnpackIntoUpdateTokenThreshold(data []byte) (*UpdateTokenThreshold, error) {
	if len(data) != 64 {
		return nil, fmt.Errorf("UpdateTokenThreshold wrong length: want 64, have %d", len(data))
	}
	return &UpdateTokenThreshold{
		Token:     common.BytesToAddress(data[:32]),
		Threshold: new(big.Int).SetBytes(data[32:64]),
	}, nil
}

func (d *UpdateTokenThreshold) requestType() byte            { return GoatUpdateTokenThresholdRequestType }
func (d *UpdateTokenThreshold) encode(b *bytes.Buffer) error { return rlp.Encode(b, d) }
func (d *UpdateTokenThreshold) decode(input []byte) error    { return rlp.DecodeBytes(input, d) }
func (d *UpdateTokenThreshold) copy() RequestData {
	return &UpdateTokenThreshold{
		Token:     d.Token,
		Threshold: new(big.Int).Set(d.Threshold),
	}
}

var (
	GoatCreateValidatorTopic      = common.HexToHash("0xf3aa84440b70359721372633122645674adb6dbb72622a222627248ef053a7dd")
	GoatValidatorLockTopic        = common.HexToHash("0xec36c0364d931187a76cf66d7eee08fad0ec2e8b7458a8d8b26b36769d4d13f3")
	GoatValidatorUnlockTopic      = common.HexToHash("0x40f2a8c5e2e2a9ad2f4e4dfc69825595b526178445c3eb22b02edfd190601db7")
	GoatClaimRewardTopic          = common.HexToHash("0xa983a6cfc4bd1095dac7b145ae020ba08e16cc7efa2051cc6b77e4011b9ee99b")
	GoatUpdateTokenWeightTopic    = common.HexToHash("0xb59bf4596e5415117fb4625044cb5b0ca5b273742825b026d06afe82a48e6217")
	GoatUpdateTokenThresholdTopic = common.HexToHash("0x326e29ab1c62c7d77fdfb302916e82e1a54f3b9961db75ee7e18afe488a0e92d")
)

func GetLockingRequests(topics []common.Hash, data []byte) (Requests, error) {
	if len(topics) != 1 {
		return nil, nil
	}

	var reqs Requests
	switch topics[0] {
	case GoatCreateValidatorTopic:
		req, err := UnpackIntoCreateValidator(data)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, NewRequest(req))
	case GoatValidatorLockTopic:
		req, err := UnpackIntoValidatorLock(data)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, NewRequest(req))
	case GoatValidatorUnlockTopic:
		req, err := UnpackIntoValidatorUnlock(data)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, NewRequest(req))
	case GoatClaimRewardTopic:
		req, err := UnpackIntoGoatRewardClaim(data)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, NewRequest(req))
	case GoatUpdateTokenWeightTopic:
		req, err := UnpackIntoUpdateTokenWeight(data)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, NewRequest(req))
	case GoatUpdateTokenThresholdTopic:
		req, err := UnpackIntoUpdateTokenThreshold(data)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, NewRequest(req))
	}

	return reqs, nil
}
