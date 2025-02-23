package transaction

import "github.com/ethereum/go-ethereum/rlp"

type EditCandidateCommissionData struct {
	PubKey     PublicKey
	Commission uint32
}

// Type returns Data type of the transaction.
func (d *EditCandidateCommissionData) Type() Type {
	return TypeEditCommissionCandidate
}

// Encode returns the byte representation of a transaction Data.
func (d *EditCandidateCommissionData) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(d)
}
