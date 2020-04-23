// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// KubernetesPod kubernetes pod
//
// swagger:model KubernetesPod
type KubernetesPod struct {

	// Time at which the pod was deployed
	DeployedTimestamp float32 `json:"deployed_timestamp,omitempty"`

	// name of the pod in Kubernetes
	Name string `json:"name,omitempty"`

	// The status of the pod
	Phase string `json:"phase,omitempty"`

	// Stdout and stderr tail of the task
	TailLines *TaskTailLines `json:"tail_lines,omitempty"`
}

// Validate validates this kubernetes pod
func (m *KubernetesPod) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTailLines(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KubernetesPod) validateTailLines(formats strfmt.Registry) error {

	if swag.IsZero(m.TailLines) { // not required
		return nil
	}

	if m.TailLines != nil {
		if err := m.TailLines.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tail_lines")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KubernetesPod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubernetesPod) UnmarshalBinary(b []byte) error {
	var res KubernetesPod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}