// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AmbariStackDetails ambari stack details
// swagger:model AmbariStackDetails

type AmbariStackDetails struct {

	// enable gpl repository
	EnableGplRepo *bool `json:"enableGplRepo,omitempty"`

	// url the MPACK that needs to be installed before HDF installation
	MpackURL string `json:"mpackUrl,omitempty"`

	// operating system for the stack, like redhat6
	Os string `json:"os,omitempty"`

	// version of the repository for VDF file creation in Ambari
	RepositoryVersion string `json:"repositoryVersion,omitempty"`

	// name of the stack, like HDP
	// Required: true
	Stack *string `json:"stack"`

	// url of the stack repository
	StackBaseURL string `json:"stackBaseURL,omitempty"`

	// id of the stack repository
	StackRepoID string `json:"stackRepoId,omitempty"`

	// url of the stack utils repository
	UtilsBaseURL string `json:"utilsBaseURL,omitempty"`

	// id of the stack utils repository
	UtilsRepoID string `json:"utilsRepoId,omitempty"`

	// whether to verify or not the repo url
	// Required: true
	Verify bool `json:"verify"`

	// version of the stack
	// Required: true
	Version *string `json:"version"`

	// local path on the Ambari server or URL that point to the desired VDF file
	VersionDefinitionFileURL string `json:"versionDefinitionFileUrl,omitempty"`
}

/* polymorph AmbariStackDetails enableGplRepo false */

/* polymorph AmbariStackDetails mpackUrl false */

/* polymorph AmbariStackDetails os false */

/* polymorph AmbariStackDetails repositoryVersion false */

/* polymorph AmbariStackDetails stack false */

/* polymorph AmbariStackDetails stackBaseURL false */

/* polymorph AmbariStackDetails stackRepoId false */

/* polymorph AmbariStackDetails utilsBaseURL false */

/* polymorph AmbariStackDetails utilsRepoId false */

/* polymorph AmbariStackDetails verify false */

/* polymorph AmbariStackDetails version false */

/* polymorph AmbariStackDetails versionDefinitionFileUrl false */

// Validate validates this ambari stack details
func (m *AmbariStackDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStack(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVerify(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AmbariStackDetails) validateStack(formats strfmt.Registry) error {

	if err := validate.Required("stack", "body", m.Stack); err != nil {
		return err
	}

	return nil
}

func (m *AmbariStackDetails) validateVerify(formats strfmt.Registry) error {

	if err := validate.Required("verify", "body", bool(m.Verify)); err != nil {
		return err
	}

	return nil
}

func (m *AmbariStackDetails) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AmbariStackDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AmbariStackDetails) UnmarshalBinary(b []byte) error {
	var res AmbariStackDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
