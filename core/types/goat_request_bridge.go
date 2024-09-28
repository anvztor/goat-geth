package types

import (
	"bytes"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

//go:generate go run github.com/fjl/gencodec -type BridgeWithdrawal -field-override bridgeWithdrawMarshaling -out gen_goat_request_withdrawal_json.go
type BridgeWithdrawal struct {
	Id         uint64 `json:"id"`
	Amount     uint64 `json:"amount_in_satoshi"`
	MaxTxPrice uint64 `json:"max_tx_price"`
	Address    string `json:"address"`
}

type bridgeWithdrawMarshaling struct {
	Id         hexutil.Uint64
	Amount     hexutil.Uint64
	MaxTxPrice hexutil.Uint64
}

var (
	withdrawalAddressLocation = big.NewInt(128)
	maxWithdrawalAddressLen   = big.NewInt(90)
	satoshi                   = big.NewInt(1e10)
)

func UnpackIntoBridgeWithdraw(topics []common.Hash, data []byte) (*BridgeWithdrawal, error) {
	if len(topics) != 3 {
		return nil, fmt.Errorf("invalid Withdraw event topics length: expect 3 got %d", len(topics))
	}

	if len(data) < 192 {
		return nil, fmt.Errorf("invalid Withdraw event data length: %d", len(data))
	}

	id := new(big.Int).SetBytes(topics[1][:])
	if !id.IsUint64() {
		return nil, fmt.Errorf("withdrawal id is too large")
	}

	amount := new(big.Int).SetBytes(data[:32]) // amount
	_, dust := amount.DivMod(amount, satoshi, new(big.Int))
	if !amount.IsUint64() {
		return nil, fmt.Errorf("withdrawal amount is too large: %d", amount)
	}

	if dust.BitLen() != 0 {
		return nil, fmt.Errorf("withdrawal amount has dust: %d", dust)
	}

	maxTxPrice := new(big.Int).SetBytes(data[64:96])
	if !maxTxPrice.IsUint64() {
		return nil, fmt.Errorf("max tx price is too large: %d", maxTxPrice)
	}

	// receiver
	if addrLoc := new(big.Int).SetBytes(data[96:128]); addrLoc.Cmp(withdrawalAddressLocation) != 0 {
		return nil, fmt.Errorf("address location should be 128 but goat %d", addrLoc)
	}

	addrLen := new(big.Int).SetBytes(data[128:160]) // length
	if addrLen.Cmp(maxWithdrawalAddressLen) > 0 {
		return nil, fmt.Errorf("address length too long: %d", addrLen)
	}

	addrLenInt64 := addrLen.Int64()
	if int64(len(data[160:])) < addrLenInt64 {
		return nil, errors.New("address slice is out of range")
	}

	return &BridgeWithdrawal{
		Id:         id.Uint64(),
		Amount:     amount.Uint64(),
		MaxTxPrice: maxTxPrice.Uint64(),
		Address:    string(data[160 : 160+addrLenInt64]),
	}, nil
}

type BridgeWithdrawals []*BridgeWithdrawal

func (s BridgeWithdrawals) Len() int { return len(s) }

func (s BridgeWithdrawals) EncodeIndex(i int, w *bytes.Buffer) {
	rlp.Encode(w, s[i])
}

// Requests creates a deep copy of each deposit and returns a slice of the
// BridgeWithdrawals requests as Request objects.
func (s BridgeWithdrawals) Requests() (reqs Requests) {
	for _, d := range s {
		reqs = append(reqs, NewRequest(d))
	}
	return
}

func (d *BridgeWithdrawal) requestType() byte            { return GoatWithdrawalRequestType }
func (d *BridgeWithdrawal) encode(b *bytes.Buffer) error { return rlp.Encode(b, d) }
func (d *BridgeWithdrawal) decode(input []byte) error    { return rlp.DecodeBytes(input, d) }
func (d *BridgeWithdrawal) copy() RequestData {
	return &BridgeWithdrawal{
		Id:         d.Id,
		Amount:     d.Amount,
		MaxTxPrice: d.MaxTxPrice,
		Address:    d.Address,
	}
}

//go:generate go run github.com/fjl/gencodec -type ReplaceByFee -field-override replaceByFeeMarshaling -out gen_goat_request_rbf_json.go
type ReplaceByFee struct {
	Id         uint64 `json:"id"`
	MaxTxPrice uint64 `json:"max_tx_price"`
}

type ReplaceByFees []*ReplaceByFee

func (s ReplaceByFees) Len() int { return len(s) }

func (s ReplaceByFees) EncodeIndex(i int, w *bytes.Buffer) {
	rlp.Encode(w, s[i])
}

// Requests creates a deep copy of each deposit and returns a slice of the
// ReplaceByFees requests as Request objects.
func (s ReplaceByFees) Requests() (reqs Requests) {
	for _, d := range s {
		reqs = append(reqs, NewRequest(d))
	}
	return
}

type replaceByFeeMarshaling struct {
	Id         hexutil.Uint64
	MaxTxPrice hexutil.Uint64
}

func UnpackIntoReplaceByFee(topics []common.Hash, data []byte) (*ReplaceByFee, error) {
	if len(topics) != 2 {
		return nil, fmt.Errorf("invalid rbf event topics length: expect 3 got %d", len(topics))
	}

	if len(data) != 32 {
		return nil, fmt.Errorf("invalid rbf event data length: %d", len(data))
	}

	id := new(big.Int).SetBytes(topics[1][:])
	if !id.IsUint64() {
		return nil, fmt.Errorf("withdrawal id is too large")
	}

	maxTxPrice := new(big.Int).SetBytes(data) // maxTxPrice
	if !maxTxPrice.IsUint64() {
		return nil, fmt.Errorf("max tx price is too large")
	}

	return &ReplaceByFee{
		Id:         id.Uint64(),
		MaxTxPrice: maxTxPrice.Uint64(),
	}, nil
}

func (d *ReplaceByFee) requestType() byte            { return GoatReplaceByFeeRequestType }
func (d *ReplaceByFee) encode(b *bytes.Buffer) error { return rlp.Encode(b, d) }
func (d *ReplaceByFee) decode(input []byte) error    { return rlp.DecodeBytes(input, d) }
func (d *ReplaceByFee) copy() RequestData {
	return &ReplaceByFee{
		Id:         d.Id,
		MaxTxPrice: d.MaxTxPrice,
	}
}

//go:generate go run github.com/fjl/gencodec -type Cancel1 -field-override cancel1Marshaling -out gen_goat_request_cancel1_json.go
type Cancel1 struct {
	Id uint64 `json:"id"`
}

type cancel1Marshaling struct {
	Id hexutil.Uint64
}

type Cancel1s []*Cancel1

func (s Cancel1s) Len() int { return len(s) }

func (s Cancel1s) EncodeIndex(i int, w *bytes.Buffer) {
	rlp.Encode(w, s[i])
}

// Requests creates a deep copy of each deposit and returns a slice of the
// Cancel1s requests as Request objects.
func (s Cancel1s) Requests() (reqs Requests) {
	for _, d := range s {
		reqs = append(reqs, NewRequest(d))
	}
	return
}

func UnpackIntoCancel1(topics []common.Hash, data []byte) (*Cancel1, error) {
	if len(topics) != 2 {
		return nil, fmt.Errorf("invalid rbf event topics length: expect 3 got %d", len(topics))
	}

	if len(data) != 0 {
		return nil, fmt.Errorf("invalid rbf event data length, expect 0 got %d", len(data))
	}

	id := new(big.Int).SetBytes(topics[1][:])
	if !id.IsUint64() {
		return nil, fmt.Errorf("withdrawal id is too large")
	}

	return &Cancel1{Id: id.Uint64()}, nil
}

func (d *Cancel1) requestType() byte            { return GoatCancel1RequestType }
func (d *Cancel1) encode(b *bytes.Buffer) error { return rlp.Encode(b, d) }
func (d *Cancel1) decode(input []byte) error    { return rlp.DecodeBytes(input, d) }
func (d *Cancel1) copy() RequestData {
	return &Cancel1{
		Id: d.Id,
	}
}

var (
	GoatAddVoterTopoic    = common.HexToHash("0x101c617f43dd1b8a54a9d747d9121bbc55e93b88bc50560d782a79c4e28fc838")
	GoatRemoveVoterTopic  = common.HexToHash("0x183393fc5cffbfc7d03d623966b85f76b9430f42d3aada2ac3f3deabc78899e8")
	GoatWithdrawalTopic   = common.HexToHash("0xbe7c38d37e8132b1d2b29509df9bf58cf1126edf2563c00db0ef3a271fb9f35b")
	GoatReplaceByFeeTopic = common.HexToHash("0x19875a7124af51c604454b74336ce2168c45bceade9d9a1e6dfae9ba7d31b7fa")
	GoatCancel1Topic      = common.HexToHash("0x0106f4416537efff55311ef5e2f9c2a48204fcf84731f2b9d5091d23fc52160c")
)

func GetBridgeRequests(topics []common.Hash, data []byte) (Requests, error) {
	if len(topics) < 2 {
		return nil, nil
	}

	var reqs Requests
	switch topics[0] {
	case GoatWithdrawalTopic:
		req, err := UnpackIntoBridgeWithdraw(topics, data)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, NewRequest(req))
	case GoatReplaceByFeeTopic:
		req, err := UnpackIntoReplaceByFee(topics, data)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, NewRequest(req))
	case GoatCancel1Topic:
		req, err := UnpackIntoCancel1(topics, data)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, NewRequest(req))
	}

	return reqs, nil
}
