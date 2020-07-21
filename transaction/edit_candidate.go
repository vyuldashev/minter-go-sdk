package transaction

import (
	"github.com/MinterTeam/minter-go-sdk/wallet"
	"github.com/ethereum/go-ethereum/rlp"
)

// Transaction for editing existing candidate.
type EditCandidateData struct {
	PubKey         [32]byte
	NewPubKey      *[32]byte `rlp:"nil"` // optional
	RewardAddress  [20]byte
	OwnerAddress   [20]byte
	ControlAddress [20]byte
}

func NewEditCandidateData() *EditCandidateData {
	return &EditCandidateData{}
}

// Set public key of a validator.
func (d *EditCandidateData) SetPubKey(key string) (*EditCandidateData, error) {
	pk, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	var pubKey [32]byte
	copy(pubKey[:], pk)
	d.PubKey = pubKey
	return d, nil
}

// Tries to set public key of validator and panics on error.
func (d *EditCandidateData) MustSetPubKey(key string) *EditCandidateData {
	_, err := d.SetPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *EditCandidateData) SetNewPubKey(key string) (*EditCandidateData, error) {
	newPubKey, err := wallet.PublicKeyToHex(key)
	if err != nil {
		return d, err
	}
	var pubKey [32]byte
	copy(pubKey[:], newPubKey)
	d.NewPubKey = &pubKey
	return d, nil
}
func (d *EditCandidateData) MustSetNewPubKey(key string) *EditCandidateData {
	_, err := d.SetNewPubKey(key)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *EditCandidateData) MustSetRewardAddress(address string) *EditCandidateData {
	_, err := d.SetRewardAddress(address)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *EditCandidateData) SetRewardAddress(address string) (*EditCandidateData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.RewardAddress[:], bytes)
	return d, nil
}

func (d *EditCandidateData) MustSetOwnerAddress(address string) *EditCandidateData {
	_, err := d.SetOwnerAddress(address)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *EditCandidateData) SetOwnerAddress(address string) (*EditCandidateData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.OwnerAddress[:], bytes)
	return d, nil
}

func (d *EditCandidateData) MustSetControlAddress(address string) *EditCandidateData {
	_, err := d.SetControlAddress(address)
	if err != nil {
		panic(err)
	}
	return d
}

func (d *EditCandidateData) SetControlAddress(address string) (*EditCandidateData, error) {
	bytes, err := wallet.AddressToHex(address)
	if err != nil {
		return d, err
	}
	copy(d.ControlAddress[:], bytes)
	return d, nil
}

func (d *EditCandidateData) Type() Type {
	return TypeEditCandidate
}

func (d *EditCandidateData) Fee() Fee {
	return feeTypeEditCandidate
}

func (d *EditCandidateData) encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
