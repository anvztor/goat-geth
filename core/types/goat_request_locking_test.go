package types

import (
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

func TestUnpackIntoCreateValidator(t *testing.T) {
	tests := []struct {
		name    string
		args    []byte
		want    *CreateValidator
		wantErr bool
	}{
		{
			name: "1",
			args: hexutil.MustDecode("0x0000000000000000000000008945a1288dc78a6d8952a92c77aee6730b4147780000000000000000000000005b38da6a701c568545dcfcb03fcb875f56beddc4b21124a8e21a475a08e4bf1ad6940f52b105d065075f610227089981948d81b0df0b6e43fc4c228a48ff159c3e6a38eb0e6ce15d78312a445d3d1671fe756842"),
			want: &CreateValidator{Pubkey: hexutil.MustDecode("0xb21124a8e21a475a08e4bf1ad6940f52b105d065075f610227089981948d81b0df0b6e43fc4c228a48ff159c3e6a38eb0e6ce15d78312a445d3d1671fe756842")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackIntoCreateValidator(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackIntoCreateValidator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackIntoCreateValidator() = %v, want %v", got, tt.want)
			}
			if got.requestType() != GoatCreateValidatorType {
				t.Errorf("UnpackIntoCreateValidator() = not GoatCreateValidatorType")
			}
			if !reflect.DeepEqual(got.copy(), got) {
				t.Errorf("UnpackIntoCreateValidator(): copy is not DeepEqual")
			}
		})
	}
}

func TestUnpackIntoValidatorLock(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *ValidatorLock
		wantErr bool
	}{
		{
			name: "1",
			args: args{hexutil.MustDecode("0x0000000000000000000000008945a1288dc78a6d8952a92c77aee6730b4147780000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a")},
			want: &ValidatorLock{
				Validator: common.HexToAddress("0x8945A1288dc78A6D8952a92C77aEe6730B414778"),
				Token:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
				Amount:    big.NewInt(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackIntoValidatorLock(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackIntoValidatorLock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackIntoValidatorLock() = %v, want %v", got, tt.want)
			}
			if got.requestType() != GoatLockType {
				t.Errorf("UnpackIntoValidatorLock() = not GoatLockType")
			}
			if !reflect.DeepEqual(got.copy(), got) {
				t.Errorf("UnpackIntoValidatorLock(): copy is not DeepEqual")
			}
		})
	}
}

func TestUnpackIntoValidatorUnlock(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *ValidatorUnlock
		wantErr bool
	}{
		{
			name: "1",
			args: args{hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000008945a1288dc78a6d8952a92c77aee6730b4147780000000000000000000000005b38da6a701c568545dcfcb03fcb875f56beddc40000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a")},
			want: &ValidatorUnlock{
				Id:        0,
				Validator: common.HexToAddress("0x8945A1288dc78A6D8952a92C77aEe6730B414778"),
				Recipient: common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				Token:     common.HexToAddress("0x0000000000000000000000000000000000000000"),
				Amount:    big.NewInt(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackIntoValidatorUnlock(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackIntoValidatorUnlock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackIntoValidatorUnlock() = %v, want %v", got, tt.want)
			}
			if got.requestType() != GoatUnlockType {
				t.Errorf("UnpackIntoValidatorUnlock() = not GoatUnlockType")
			}
			if !reflect.DeepEqual(got.copy(), got) {
				t.Errorf("UnpackIntoValidatorUnlock(): copy is not DeepEqual")
			}
		})
	}
}

func TestUnpackIntoGoatRewardClaim(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *GoatRewardClaim
		wantErr bool
	}{
		{
			name: "1",
			args: args{hexutil.MustDecode("0x00000000000000000000000000000000000000000000000000000000000000010000000000000000000000008945a1288dc78a6d8952a92c77aee6730b4147780000000000000000000000005b38da6a701c568545dcfcb03fcb875f56beddc4")},
			want: &GoatRewardClaim{
				Id:        1,
				Validator: common.HexToAddress("0x8945A1288dc78A6D8952a92C77aEe6730B414778"),
				Recipient: common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackIntoGoatRewardClaim(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackIntoGoatRewardClaim() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackIntoGoatRewardClaim() = %v, want %v", got, tt.want)
			}
			if got.requestType() != GoatClaimRewardType {
				t.Errorf("UnpackIntoGoatRewardClaim() = not GoatClaimType")
			}
			if !reflect.DeepEqual(got.copy(), got) {
				t.Errorf("UnpackIntoGoatRewardClaim(): copy is not DeepEqual")
			}
		})
	}
}

func TestUnpackIntoSetTokenWeight(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *SetTokenWeight
		wantErr bool
	}{
		{
			name: "1",
			args: args{hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000a")},
			want: &SetTokenWeight{
				Token:  common.HexToAddress("0x0000000000000000000000000000000000000000"),
				Weight: 10,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackIntoSetTokenWeight(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackIntoSetTokenWeight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackIntoSetTokenWeight() = %v, want %v", got, tt.want)
			}
			if got.requestType() != GoatSetTokenWeight {
				t.Errorf("UnpackIntoSetTokenWeight() = not GoatSetTokenWeight")
			}
			if !reflect.DeepEqual(got.copy(), got) {
				t.Errorf("UnpackIntoSetTokenWeight(): copy is not DeepEqual")
			}
		})
	}
}

func TestUnpackIntoSetTokenThreshold(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *SetTokenThreshold
		wantErr bool
	}{
		{
			name: "1",
			args: args{hexutil.MustDecode("0x0000000000000000000000005b38da6a701c568545dcfcb03fcb875f56beddc4000000000000000000000000000000000000000000000000000000000000000a")},
			want: &SetTokenThreshold{
				Token:     common.HexToAddress("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4"),
				Threshold: big.NewInt(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackIntoSetTokenThreshold(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackIntoSetTokenThreshold() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackIntoSetTokenThreshold() = %v, want %v", got, tt.want)
			}
			if got.requestType() != GoatSetTokenThreshold {
				t.Errorf("UnpackIntoSetTokenThreshold() = not GoatSetTokenThreshold")
			}
			if !reflect.DeepEqual(got.copy(), got) {
				t.Errorf("UnpackIntoSetTokenThreshold(): copy is not DeepEqual")
			}
		})
	}
}
