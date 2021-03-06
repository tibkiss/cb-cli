// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AmbariStackDetailsResponse ambari stack details response
// swagger:model AmbariStackDetailsResponse

type AmbariStackDetailsResponse struct {

	// enable gpl repo
	EnableGplRepo *bool `json:"enableGplRepo,omitempty"`

	// hdp version
	HdpVersion string `json:"hdpVersion,omitempty"`

	// knox
	Knox map[string]string `json:"knox,omitempty"`

	// stack
	Stack map[string]string `json:"stack,omitempty"`

	// util
	Util map[string]string `json:"util,omitempty"`

	// verify
	Verify *bool `json:"verify,omitempty"`
}

/* polymorph AmbariStackDetailsResponse enableGplRepo false */

/* polymorph AmbariStackDetailsResponse hdpVersion false */

/* polymorph AmbariStackDetailsResponse knox false */

/* polymorph AmbariStackDetailsResponse stack false */

/* polymorph AmbariStackDetailsResponse util false */

/* polymorph AmbariStackDetailsResponse verify false */

// Validate validates this ambari stack details response
func (m *AmbariStackDetailsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AmbariStackDetailsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AmbariStackDetailsResponse) UnmarshalBinary(b []byte) error {
	var res AmbariStackDetailsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
