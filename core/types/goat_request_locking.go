package types

import (
	"bytes"
	"math/big"

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

func GetLockingRequests(topics []common.Hash, data []byte) (Requests, error) {
	return nil, nil
}
