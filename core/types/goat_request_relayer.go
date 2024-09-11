package types

import (
	"bytes"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rlp"
)

type AddVoters []*AddVoter

func (s AddVoters) Len() int { return len(s) }

func (s AddVoters) EncodeIndex(i int, w *bytes.Buffer) {
	rlp.Encode(w, s[i])
}

// Requests creates a deep copy of each deposit and returns a slice of the
// AddVoters requests as Request objects.
func (s AddVoters) Requests() (reqs Requests) {
	for _, d := range s {
		reqs = append(reqs, NewRequest(d))
	}
	return
}

type AddVoter struct {
	Voter  common.Address `json:"voter"`
	Pubkey common.Hash    `json:"pubkey"`
}

func (d *AddVoter) requestType() byte            { return GoatAddVoterRequestType }
func (d *AddVoter) encode(b *bytes.Buffer) error { return rlp.Encode(b, d) }
func (d *AddVoter) decode(input []byte) error    { return rlp.DecodeBytes(input, d) }
func (d *AddVoter) copy() RequestData {
	return &AddVoter{
		Voter:  d.Voter,
		Pubkey: d.Pubkey,
	}
}

func UnpackIntoAddVoter(topics []common.Hash, data []byte) (*AddVoter, error) {
	if len(topics) != 2 {
		return nil, fmt.Errorf("invalid AddVoter event topics length: expect 2 got %d", len(topics))
	}

	if len(data) != 32 {
		return nil, fmt.Errorf("addVoter wrong length: want 64, have %d", len(data))
	}

	return &AddVoter{
		Voter:  common.BytesToAddress(topics[1][:]),
		Pubkey: common.BytesToHash(data[:]),
	}, nil
}

type RemoveVoters []*RemoveVoter

func (s RemoveVoters) Len() int { return len(s) }

func (s RemoveVoters) EncodeIndex(i int, w *bytes.Buffer) {
	rlp.Encode(w, s[i])
}

// Requests creates a deep copy of each deposit and returns a slice of the
// RemoveVoters requests as Request objects.
func (s RemoveVoters) Requests() (reqs Requests) {
	for _, d := range s {
		reqs = append(reqs, NewRequest(d))
	}
	return
}

type RemoveVoter struct {
	Voter common.Address `json:"voter"`
}

func (d *RemoveVoter) requestType() byte            { return GoatRemoveVoterRequestType }
func (d *RemoveVoter) encode(b *bytes.Buffer) error { return rlp.Encode(b, d) }
func (d *RemoveVoter) decode(input []byte) error    { return rlp.DecodeBytes(input, d) }
func (d *RemoveVoter) copy() RequestData {
	return &RemoveVoter{
		Voter: d.Voter,
	}
}

func UnpackIntoRemoveVoter(topics []common.Hash, data []byte) (*RemoveVoter, error) {
	if len(topics) != 2 {
		return nil, fmt.Errorf("invalid RemoveVoter event topics length: expect 2 got %d", len(topics))
	}
	if len(data) != 0 {
		return nil, fmt.Errorf("RemoveVoter wrong length: want 0, have %d", len(data))
	}
	return &RemoveVoter{Voter: common.BytesToAddress(topics[1][:])}, nil
}

func GetRelayerRequests(topics []common.Hash, data []byte) (Requests, error) {
	if len(topics) != 2 {
		return nil, nil
	}

	var reqs Requests
	switch topics[0] {
	case GoatAddVoterTopoic:
		req, err := UnpackIntoAddVoter(topics, data)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, NewRequest(req))
	case GoatRemoveVoterTopic:
		req, err := UnpackIntoRemoveVoter(topics, data)
		if err != nil {
			return nil, err
		}
		reqs = append(reqs, NewRequest(req))
	}

	return reqs, nil
}
