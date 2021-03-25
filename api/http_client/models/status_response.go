// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StatusResponse status response
//
// swagger:model StatusResponse
type StatusResponse struct {

	// catching up
	CatchingUp bool `json:"catching_up,omitempty"`

	// initial height
	InitialHeight uint64 `json:"initial_height,omitempty,string"`

	// keep last states
	KeepLastStates uint64 `json:"keep_last_states,omitempty,string"`

	// latest app hash
	LatestAppHash string `json:"latest_app_hash,omitempty"`

	// latest block hash
	LatestBlockHash string `json:"latest_block_hash,omitempty"`

	// latest block height
	LatestBlockHeight uint64 `json:"latest_block_height,omitempty,string"`

	// latest block time
	LatestBlockTime string `json:"latest_block_time,omitempty"`

	// network
	Network string `json:"network,omitempty"`

	// node id
	NodeID string `json:"node_id,omitempty"`

	// public key
	PublicKey string `json:"public_key,omitempty"`

	// total slashed
	TotalSlashed string `json:"total_slashed,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this status response
func (m *StatusResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusResponse) UnmarshalBinary(b []byte) error {
	var res StatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
