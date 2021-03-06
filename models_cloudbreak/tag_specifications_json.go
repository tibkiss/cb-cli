// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// TagSpecificationsJSON tag specifications Json
// swagger:model TagSpecificationsJson

type TagSpecificationsJSON struct {

	// tag specifications
	Specifications map[string]map[string]interface{} `json:"specifications,omitempty"`
}

/* polymorph TagSpecificationsJson specifications false */

// Validate validates this tag specifications Json
func (m *TagSpecificationsJSON) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *TagSpecificationsJSON) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagSpecificationsJSON) UnmarshalBinary(b []byte) error {
	var res TagSpecificationsJSON
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
