// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateIdentityRequestPayload update identity request payload
//
// swagger:model UpdateIdentityRequestPayload
type UpdateIdentityRequestPayload struct {

	// SchemaID is the ID of the JSON Schema to be used for validating the identity's traits. If set
	// will update the Identity's SchemaID.
	SchemaID string `json:"schema_id,omitempty"`

	// Traits represent an identity's traits. The identity is able to create, modify, and delete traits
	// in a self-service manner. The input will always be validated against the JSON Schema defined
	// in `schema_id`.
	// Required: true
	Traits interface{} `json:"traits"`
}

// Validate validates this update identity request payload
func (m *UpdateIdentityRequestPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTraits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateIdentityRequestPayload) validateTraits(formats strfmt.Registry) error {

	if err := validate.Required("traits", "body", m.Traits); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateIdentityRequestPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateIdentityRequestPayload) UnmarshalBinary(b []byte) error {
	var res UpdateIdentityRequestPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}