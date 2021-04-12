// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// CandidatesRequestCandidateStatus candidates request candidate status
//
// swagger:model CandidatesRequestCandidateStatus
type CandidatesRequestCandidateStatus string

const (

	// CandidatesRequestCandidateStatusAll captures enum value "all"
	CandidatesRequestCandidateStatusAll CandidatesRequestCandidateStatus = "all"

	// CandidatesRequestCandidateStatusOff captures enum value "off"
	CandidatesRequestCandidateStatusOff CandidatesRequestCandidateStatus = "off"

	// CandidatesRequestCandidateStatusOn captures enum value "on"
	CandidatesRequestCandidateStatusOn CandidatesRequestCandidateStatus = "on"

	// CandidatesRequestCandidateStatusValidator captures enum value "validator"
	CandidatesRequestCandidateStatusValidator CandidatesRequestCandidateStatus = "validator"
)

// for schema
var candidatesRequestCandidateStatusEnum []interface{}

func init() {
	var res []CandidatesRequestCandidateStatus
	if err := json.Unmarshal([]byte(`["all","off","on","validator"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		candidatesRequestCandidateStatusEnum = append(candidatesRequestCandidateStatusEnum, v)
	}
}

func (m CandidatesRequestCandidateStatus) validateCandidatesRequestCandidateStatusEnum(path, location string, value CandidatesRequestCandidateStatus) error {
	if err := validate.EnumCase(path, location, value, candidatesRequestCandidateStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this candidates request candidate status
func (m CandidatesRequestCandidateStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCandidatesRequestCandidateStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
