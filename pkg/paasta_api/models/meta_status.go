// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetaStatus Result of `paasta metastatus` command
//
// swagger:model MetaStatus
type MetaStatus struct {

	// Exit code from `paasta metastatus` command
	ExitCode int64 `json:"exit_code,omitempty"`

	// Output from `paasta metastatus` command
	Output string `json:"output,omitempty"`
}

// Validate validates this meta status
func (m *MetaStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetaStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetaStatus) UnmarshalBinary(b []byte) error {
	var res MetaStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}