package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*StackValidationRequest stack validation request

swagger:model StackValidationRequest
*/
type StackValidationRequest struct {

	/* id of the referenced blueprint

	Required: true
	*/
	BlueprintID int64 `json:"blueprintId"`

	/* file system
	 */
	FileSystem *FileSystem `json:"fileSystem,omitempty"`

	/* host groups

	Required: true
	Unique: true
	*/
	HostGroups []*HostGroup `json:"hostGroups"`

	/* instance groups

	Required: true
	Unique: true
	*/
	InstanceGroups []*InstanceGroup `json:"instanceGroups"`

	/* network resource id for the stack

	Required: true
	*/
	NetworkID int64 `json:"networkId"`

	/* type of cloud provider

	Required: true
	*/
	Platform string `json:"platform"`
}

// Validate validates this stack validation request
func (m *StackValidationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlueprintID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHostGroups(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInstanceGroups(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetworkID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePlatform(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackValidationRequest) validateBlueprintID(formats strfmt.Registry) error {

	if err := validate.Required("blueprintId", "body", int64(m.BlueprintID)); err != nil {
		return err
	}

	return nil
}

func (m *StackValidationRequest) validateHostGroups(formats strfmt.Registry) error {

	if err := validate.Required("hostGroups", "body", m.HostGroups); err != nil {
		return err
	}

	if err := validate.UniqueItems("hostGroups", "body", m.HostGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.HostGroups); i++ {

		if m.HostGroups[i] != nil {

			if err := m.HostGroups[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *StackValidationRequest) validateInstanceGroups(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroups", "body", m.InstanceGroups); err != nil {
		return err
	}

	if err := validate.UniqueItems("instanceGroups", "body", m.InstanceGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.InstanceGroups); i++ {

		if m.InstanceGroups[i] != nil {

			if err := m.InstanceGroups[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *StackValidationRequest) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.Required("networkId", "body", int64(m.NetworkID)); err != nil {
		return err
	}

	return nil
}

func (m *StackValidationRequest) validatePlatform(formats strfmt.Registry) error {

	if err := validate.RequiredString("platform", "body", string(m.Platform)); err != nil {
		return err
	}

	return nil
}