package types

import (
	"math/big"
	"reflect"
	"testing"
)

func TestNewGoatGasRevenue(t *testing.T) {
	type args struct {
		amount *big.Int
	}
	tests := []struct {
		name string
		args args
		want *GasRevenue
	}{
		{"1", args{big.NewInt(100)}, &GasRevenue{big.NewInt(100)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGoatGasRevenue(tt.args.amount)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGoatGasRevenue() = %v, want %v", got, tt.want)
			}

			if got.requestType() != GoatGasRevenueRequestType {
				t.Errorf("NewGoatGasRevenue() = not GoatGasRevenueRequestType")
			}

			if !reflect.DeepEqual(got.copy(), got) {
				t.Errorf("NewGoatGasRevenue(): copy is not DeepEqual")
			}
		})
	}
}
