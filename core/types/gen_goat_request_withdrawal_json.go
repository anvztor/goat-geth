// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*goatWithdrawMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (g GoatWithdrawal) MarshalJSON() ([]byte, error) {
	type GoatWithdrawal struct {
		Id         hexutil.Uint64 `json:"id"`
		Amount     hexutil.Uint64 `json:"amount_in_satoshi"`
		MaxTxPrice hexutil.Uint64 `json:"max_tx_price"`
		Address    string         `json:"address"`
	}
	var enc GoatWithdrawal
	enc.Id = hexutil.Uint64(g.Id)
	enc.Amount = hexutil.Uint64(g.Amount)
	enc.MaxTxPrice = hexutil.Uint64(g.MaxTxPrice)
	enc.Address = g.Address
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (g *GoatWithdrawal) UnmarshalJSON(input []byte) error {
	type GoatWithdrawal struct {
		Id         *hexutil.Uint64 `json:"id"`
		Amount     *hexutil.Uint64 `json:"amount_in_satoshi"`
		MaxTxPrice *hexutil.Uint64 `json:"max_tx_price"`
		Address    *string         `json:"address"`
	}
	var dec GoatWithdrawal
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Id != nil {
		g.Id = uint64(*dec.Id)
	}
	if dec.Amount != nil {
		g.Amount = uint64(*dec.Amount)
	}
	if dec.MaxTxPrice != nil {
		g.MaxTxPrice = uint64(*dec.MaxTxPrice)
	}
	if dec.Address != nil {
		g.Address = *dec.Address
	}
	return nil
}
