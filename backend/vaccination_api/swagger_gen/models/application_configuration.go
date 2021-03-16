// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationConfiguration application configuration
//
// swagger:model ApplicationConfiguration
type ApplicationConfiguration struct {

	// navigation
	Navigation interface{} `json:"navigation,omitempty"`

	// styles
	Styles interface{} `json:"styles,omitempty"`

	// validation
	Validation interface{} `json:"validation,omitempty"`
}

// Validate validates this application configuration
func (m *ApplicationConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this application configuration based on context it is used
func (m *ApplicationConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationConfiguration) UnmarshalBinary(b []byte) error {
	var res ApplicationConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
