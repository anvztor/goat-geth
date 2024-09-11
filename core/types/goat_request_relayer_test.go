package types

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func TestUnpackIntoRemoveVoter(t *testing.T) {
	type args struct {
		topics []common.Hash
		data   []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *RemoveVoter
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				topics: []common.Hash{
					common.HexToHash("0x183393fc5cffbfc7d03d623966b85f76b9430f42d3aada2ac3f3deabc78899e8"),
					common.HexToHash("0x000000000000000000000000c96397756df86d3ac4c04958ee5bf9ac7421e328"),
				},
				data: nil,
			},
			want: &RemoveVoter{
				Voter: common.HexToAddress("0xc96397756df86d3ac4c04958ee5bf9ac7421e328"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackIntoRemoveVoter(tt.args.topics, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackIntoRemoveVoter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackIntoRemoveVoter() = %v, want %v", got, tt.want)
			}

			if tt.wantErr {
				return
			}

			reqs, err := GetRelayerRequests(tt.args.topics, tt.args.data)
			if err != nil {
				t.Errorf("UnpackIntoRemoveVoter(): GetRelayerRequests error = %v", err)
				return
			}

			if len(reqs) != 1 {
				t.Errorf("UnpackIntoRemoveVoter(): GetRelayerRequests length(1) != %d", len(reqs))
				return
			}

			ty, ok := reqs[0].inner.(*RemoveVoter)
			if !ok {
				t.Errorf("UnpackIntoRemoveVoter(): GetRelayerRequests not RemoveVoter")
				return
			}
			if !reflect.DeepEqual(got, ty) {
				t.Errorf("UnpackIntoRemoveVoter() = %v, want %v", got, tt.want)
			}

			if ty.requestType() != GoatRemoveVoterRequestType {
				t.Errorf("UnpackIntoRemoveVoter() = not GoatRemoveVoterRequestType")
			}

			if !reflect.DeepEqual(ty.copy(), ty) {
				t.Errorf("UnpackIntoRemoveVoter(): copy is not DeepEqual")
			}
		})
	}
}

func TestUnpackIntoAddVoter(t *testing.T) {
	type args struct {
		topics []common.Hash
		data   []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *AddVoter
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				topics: []common.Hash{
					common.HexToHash("0x101c617f43dd1b8a54a9d747d9121bbc55e93b88bc50560d782a79c4e28fc838"),
					common.HexToHash("0x000000000000000000000000d12a5a92d4621fbe3068914988d538c410245443"),
				},
				data: hexutil.MustDecode("0x023504e3cadac49656b8f0ac939b1665870c5eb60cd47541e401babb7ff99f23"),
			},
			want: &AddVoter{
				Voter:  common.HexToAddress("0xd12a5a92D4621fBE3068914988D538c410245443"),
				Pubkey: common.HexToHash("0x023504e3cadac49656b8f0ac939b1665870c5eb60cd47541e401babb7ff99f23"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnpackIntoAddVoter(tt.args.topics, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnpackIntoAddVoter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnpackIntoAddVoter() = %v, want %v", got, tt.want)
			}

			if tt.wantErr {
				return
			}

			reqs, err := GetRelayerRequests(tt.args.topics, tt.args.data)
			if err != nil {
				t.Errorf("UnpackIntoAddVoter(): GetRelayerRequests error = %v", err)
				return
			}

			if len(reqs) != 1 {
				t.Errorf("UnpackIntoAddVoter(): GetRelayerRequests length(1) != %d", len(reqs))
				return
			}

			ty, ok := reqs[0].inner.(*AddVoter)
			if !ok {
				t.Errorf("UnpackIntoAddVoter(): GetRelayerRequests not AddVoter")
				return
			}
			if !reflect.DeepEqual(got, ty) {
				t.Errorf("UnpackIntoAddVoter() = %v, want %v", got, tt.want)
			}

			if ty.requestType() != GoatAddVoterRequestType {
				t.Errorf("UnpackIntoAddVoter() = not GoatAddVoterRequestType")
			}

			if !reflect.DeepEqual(ty.copy(), ty) {
				t.Errorf("UnpackIntoAddVoter(): copy is not DeepEqual")
			}
		})
	}
}
