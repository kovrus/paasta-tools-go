// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourceItem resource item
//
// swagger:model ResourceItem
type ResourceItem struct {

	// cpus
	Cpus *ResourceValue `json:"cpus,omitempty"`

	// disk
	Disk *ResourceValue `json:"disk,omitempty"`

	// groupings
	Groupings interface{} `json:"groupings,omitempty"`

	// mem
	Mem *ResourceValue `json:"mem,omitempty"`
}

// Validate validates this resource item
func (m *ResourceItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCpus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourceItem) validateCpus(formats strfmt.Registry) error {

	if swag.IsZero(m.Cpus) { // not required
		return nil
	}

	if m.Cpus != nil {
		if err := m.Cpus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cpus")
			}
			return err
		}
	}

	return nil
}

func (m *ResourceItem) validateDisk(formats strfmt.Registry) error {

	if swag.IsZero(m.Disk) { // not required
		return nil
	}

	if m.Disk != nil {
		if err := m.Disk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("disk")
			}
			return err
		}
	}

	return nil
}

func (m *ResourceItem) validateMem(formats strfmt.Registry) error {

	if swag.IsZero(m.Mem) { // not required
		return nil
	}

	if m.Mem != nil {
		if err := m.Mem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceItem) UnmarshalBinary(b []byte) error {
	var res ResourceItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}