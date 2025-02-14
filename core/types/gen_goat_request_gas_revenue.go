// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*gasRevenueMarshaling)(nil)

// MarshalJSON marshals as JSON.
func (g GasRevenue) MarshalJSON() ([]byte, error) {
	type GasRevenue struct {
		Amount *hexutil.Big `json:"amount"`
	}
	var enc GasRevenue
	enc.Amount = (*hexutil.Big)(g.Amount)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (g *GasRevenue) UnmarshalJSON(input []byte) error {
	type GasRevenue struct {
		Amount *hexutil.Big `json:"amount"`
	}
	var dec GasRevenue
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Amount != nil {
		g.Amount = (*big.Int)(dec.Amount)
	}
	return nil
}
