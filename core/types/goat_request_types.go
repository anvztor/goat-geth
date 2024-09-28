package types

import (
	"encoding/json"
	"errors"
)

const (
	GoatGasRevenueRequestType byte = iota + 0x60
	GoatAddVoterRequestType
	GoatRemoveVoterRequestType
	GoatWithdrawalRequestType
	GoatReplaceByFeeRequestType
	GoatCancel1RequestType
	GoatCreateValidatorType
	GoatLockType
	GoatUnlockType
	GoatClaimRewardType
	GoatSetTokenWeight
	GoatSetTokenThreshold
)

func (r *Request) ForGoat() bool {
	t := r.inner.requestType()
	return t >= GoatGasRevenueRequestType && t <= GoatSetTokenThreshold
}

// Request intermediate type for json codec
type requestMarshaling struct {
	Type byte            `json:"type"`
	Data json.RawMessage `json:"data"`
}

// UnmarshalJSON implements json.Marshaler interface
func (r *Request) MarshalJSON() ([]byte, error) {
	if r.inner == nil {
		return nil, errors.New("no request data")
	}

	data, err := json.Marshal(r.inner)
	if err != nil {
		return nil, err
	}

	return json.Marshal(requestMarshaling{Type: r.inner.requestType(), Data: data})
}

// UnmarshalJSON implements json.Unmarshaler interface
func (r *Request) UnmarshalJSON(b []byte) error {
	var raw requestMarshaling
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	switch raw.Type {
	case GoatGasRevenueRequestType:
		r.inner = new(GasRevenue)
	case GoatAddVoterRequestType:
		r.inner = new(AddVoter)
	case GoatRemoveVoterRequestType:
		r.inner = new(RemoveVoter)
	case GoatWithdrawalRequestType:
		r.inner = new(BridgeWithdrawal)
	case GoatReplaceByFeeRequestType:
		r.inner = new(ReplaceByFee)
	case GoatCancel1RequestType:
		r.inner = new(Cancel1)
	case GoatCreateValidatorType:
		r.inner = new(CreateValidator)
	case GoatLockType:
		r.inner = new(ValidatorLock)
	case GoatUnlockType:
		r.inner = new(ValidatorUnlock)
	case GoatClaimRewardType:
		r.inner = new(GoatRewardClaim)
	case GoatSetTokenWeight:
		r.inner = new(SetTokenWeight)
	case GoatSetTokenThreshold:
		r.inner = new(SetTokenThreshold)

	case DepositRequestType:
		r.inner = new(Deposit)
	default:
		return ErrRequestTypeNotSupported
	}

	return json.Unmarshal(raw.Data, r.inner)
}
