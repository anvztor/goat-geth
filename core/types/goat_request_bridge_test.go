package types

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestUnpackIntoBridgeWithdraw(t *testing.T) {
	type args struct {
		topics []common.Hash
		data   []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *BridgeWithdrawal
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				topics: []common.Hash{
					common.HexToHash("0xbe7c38d37e8132b1d2b29509df9bf58cf1126edf2563c00db0ef3a271fb9f35b"),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000064"),
					common.HexToHash("0x0000000000000000000000005b38da6a701c568545dcfcb03fcb875f56beddc4"),
				},
				data: hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000174876e800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000003e62633171656e356b76336330657064397966717675327130353971736a7077753968646a797778327639703570396c386d73786e383866733979356b78360000"),
			},
			want: &BridgeWithdrawal{
				Id:         100,
				Amount:     10,
				MaxTxPrice: 1,
				Address:    "bc1qen5kv3c0epd9yfqvu2q059qsjpwu9hdjywx2v9p5p9l8msxn88fs9y5kx6",
			},
		},
		{
			name: "2",
			args: args{
				topics: []common.Hash{
					common.HexToHash("0xbe7c38d37e8132b1d2b29509df9bf58cf1126edf2563c00db0ef3a271fb9f35b"),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
					common.HexToHash("0x0000000000000000000000005b38da6a701c568545dcfcb03fcb875f56beddc4"),
				},
				data: hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000002e90edd00000000000000000000000000000000000000000000000000000000000000003e8000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000002a626331716d76733230387765336a67376867637a686c683765397566773033346b666d3276777376676500000000000000000000000000000000000000000000"),
			},
			want: &BridgeWithdrawal{
				Id:         1,
				Amount:     20,
				MaxTxPrice: 10,
				Address:    "bc1qmvs208we3jg7hgczhlh7e9ufw034kfm2vwsvge",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackIntoBridgeWithdraw(tt.args.topics, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackIntoBridgeWithdraw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackIntoBridgeWithdraw() = %v, want %v", got, tt.want)
			}

			if tt.wantErr {
				return
			}

			reqs, err := GetBridgeRequests(tt.args.topics, tt.args.data)
			if err != nil {
				t.Errorf("UnpackIntoBridgeWithdraw(): GetBridgeRequests error = %v", err)
				return
			}

			if len(reqs) != 1 {
				t.Errorf("UnpackIntoBridgeWithdraw(): GetBridgeRequests length(1) != %d", len(reqs))
				return
			}

			ty, ok := reqs[0].inner.(*BridgeWithdrawal)
			if !ok {
				t.Errorf("UnpackIntoBridgeWithdraw(): GetBridgeRequests not BridgeWithdrawal")
				return
			}
			if !reflect.DeepEqual(got, ty) {
				t.Errorf("UnpackIntoBridgeWithdraw() = %v, want %v", got, tt.want)
			}

			if ty.requestType() != GoatWithdrawalRequestType {
				t.Errorf("UnpackIntoBridgeWithdraw() = not GoatWithdrawalRequestType")
			}

			if !reflect.DeepEqual(ty.copy(), ty) {
				t.Errorf("UnpackIntoBridgeWithdraw(): copy is not DeepEqual")
			}
		})
	}
}

func TestUnpackIntoReplaceByFee(t *testing.T) {
	type args struct {
		topics []common.Hash
		data   []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *ReplaceByFee
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				topics: []common.Hash{
					common.HexToHash("0x19875a7124af51c604454b74336ce2168c45bceade9d9a1e6dfae9ba7d31b7fa"),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
				},
				data: hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000014"),
			},
			want: &ReplaceByFee{
				Id:         1,
				MaxTxPrice: 20,
			},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				topics: []common.Hash{
					common.HexToHash("0x19875a7124af51c604454b74336ce2168c45bceade9d9a1e6dfae9ba7d31b7fa"),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"),
				},
				data: hexutil.MustDecode("0x000000000000000000000000000000000000000000000000000000000000000a"),
			},
			want: &ReplaceByFee{
				Id:         2,
				MaxTxPrice: 10,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackIntoReplaceByFee(tt.args.topics, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackIntoReplaceByFee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackIntoReplaceByFee() = %v, want %v", got, tt.want)
			}

			if tt.wantErr {
				return
			}

			reqs, err := GetBridgeRequests(tt.args.topics, tt.args.data)
			if err != nil {
				t.Errorf("UnpackIntoReplaceByFee(): GetBridgeRequests error = %v", err)
				return
			}

			if len(reqs) != 1 {
				t.Errorf("UnpackIntoReplaceByFee(): GetBridgeRequests length(1) != %d", len(reqs))
				return
			}

			ty, ok := reqs[0].inner.(*ReplaceByFee)
			if !ok {
				t.Errorf("UnpackIntoReplaceByFee(): GetBridgeRequests not ReplaceByFee")
				return
			}
			if !reflect.DeepEqual(got, ty) {
				t.Errorf("UnpackIntoReplaceByFee() = %v, want %v", got, tt.want)
			}

			if ty.requestType() != GoatReplaceByFeeRequestType {
				t.Errorf("UnpackIntoReplaceByFee() = not GoatReplaceByFeeRequestType")
			}

			if !reflect.DeepEqual(ty.copy(), ty) {
				t.Errorf("UnpackIntoReplaceByFee(): copy is not DeepEqual")
			}
		})
	}
}

func TestUnpackIntoCancel1(t *testing.T) {
	type args struct {
		topics []common.Hash
		data   []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *Cancel1
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				topics: []common.Hash{
					common.HexToHash("0x0106f4416537efff55311ef5e2f9c2a48204fcf84731f2b9d5091d23fc52160c"),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000001"),
				},
			},
			want: &Cancel1{
				Id: 1,
			},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				topics: []common.Hash{
					common.HexToHash("0x0106f4416537efff55311ef5e2f9c2a48204fcf84731f2b9d5091d23fc52160c"),
					common.HexToHash("0x0000000000000000000000000000000000000000000000000000000000000002"),
				},
			},
			want: &Cancel1{
				Id: 2,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackIntoCancel1(tt.args.topics, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackIntoCancel1() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackIntoCancel1() = %v, want %v", got, tt.want)
			}

			if tt.wantErr {
				return
			}

			reqs, err := GetBridgeRequests(tt.args.topics, tt.args.data)
			if err != nil {
				t.Errorf("UnpackIntoCancel1(): GetBridgeRequests error = %v", err)
				return
			}

			if len(reqs) != 1 {
				t.Errorf("UnpackIntoCancel1(): GetBridgeRequests length(1) != %d", len(reqs))
				return
			}

			ty, ok := reqs[0].inner.(*Cancel1)
			if !ok {
				t.Errorf("UnpackIntoCancel1(): GetBridgeRequests not Cancel1")
				return
			}
			if !reflect.DeepEqual(got, ty) {
				t.Errorf("UnpackIntoCancel1() = %v, want %v", got, tt.want)
			}

			if ty.requestType() != GoatCancel1RequestType {
				t.Errorf("UnpackIntoCancel1() = not GoatCancel1RequestType")
			}

			if !reflect.DeepEqual(ty.copy(), ty) {
				t.Errorf("UnpackIntoCancel1(): copy is not DeepEqual")
			}
		})
	}
}
