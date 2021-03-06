// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StackV2Request stack v2 request
// swagger:model StackV2Request

type StackV2Request struct {

	// specific version of ambari
	AmbariVersion string `json:"ambariVersion,omitempty"`

	// cluster request object on stack
	Cluster *ClusterV2Request `json:"cluster,omitempty"`

	// settings related to custom domain names
	CustomDomain *CustomDomainSettings `json:"customDomain,omitempty"`

	// failure policy in case of failures
	FailurePolicy *FailurePolicyRequest `json:"failurePolicy,omitempty"`

	// id of the related flex subscription
	FlexID int64 `json:"flexId,omitempty"`

	// general configuration parameters for a cluster (e.g. 'name', 'credentialname')
	// Required: true
	General *GeneralSettings `json:"general"`

	// specific version of HDP
	HdpVersion string `json:"hdpVersion,omitempty"`

	// settings for custom images
	ImageSettings *ImageSettings `json:"imageSettings,omitempty"`

	// collection of instance groupst
	// Required: true
	InstanceGroups []*InstanceGroupsV2 `json:"instanceGroups"`

	// stack related network
	Network *NetworkV2Request `json:"network,omitempty"`

	// additional cloud specific parameters for stack
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// placement configuration parameters for a cluster (e.g. 'region', 'availabilityZone')
	Placement *PlacementSettings `json:"placement,omitempty"`

	// cloud provider api variant
	PlatformVariant string `json:"platformVariant,omitempty"`

	// stack related authentication
	StackAuthentication *StackAuthentication `json:"stackAuthentication,omitempty"`

	// stack related tags
	Tags *Tags `json:"tags,omitempty"`
}

/* polymorph StackV2Request ambariVersion false */

/* polymorph StackV2Request cluster false */

/* polymorph StackV2Request customDomain false */

/* polymorph StackV2Request failurePolicy false */

/* polymorph StackV2Request flexId false */

/* polymorph StackV2Request general false */

/* polymorph StackV2Request hdpVersion false */

/* polymorph StackV2Request imageSettings false */

/* polymorph StackV2Request instanceGroups false */

/* polymorph StackV2Request network false */

/* polymorph StackV2Request parameters false */

/* polymorph StackV2Request placement false */

/* polymorph StackV2Request platformVariant false */

/* polymorph StackV2Request stackAuthentication false */

/* polymorph StackV2Request tags false */

// Validate validates this stack v2 request
func (m *StackV2Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCustomDomain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFailurePolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGeneral(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImageSettings(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInstanceGroups(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePlacement(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStackAuthentication(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackV2Request) validateCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {

		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *StackV2Request) validateCustomDomain(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomDomain) { // not required
		return nil
	}

	if m.CustomDomain != nil {

		if err := m.CustomDomain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customDomain")
			}
			return err
		}
	}

	return nil
}

func (m *StackV2Request) validateFailurePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.FailurePolicy) { // not required
		return nil
	}

	if m.FailurePolicy != nil {

		if err := m.FailurePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("failurePolicy")
			}
			return err
		}
	}

	return nil
}

func (m *StackV2Request) validateGeneral(formats strfmt.Registry) error {

	if err := validate.Required("general", "body", m.General); err != nil {
		return err
	}

	if m.General != nil {

		if err := m.General.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("general")
			}
			return err
		}
	}

	return nil
}

func (m *StackV2Request) validateImageSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageSettings) { // not required
		return nil
	}

	if m.ImageSettings != nil {

		if err := m.ImageSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("imageSettings")
			}
			return err
		}
	}

	return nil
}

func (m *StackV2Request) validateInstanceGroups(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroups", "body", m.InstanceGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.InstanceGroups); i++ {

		if swag.IsZero(m.InstanceGroups[i]) { // not required
			continue
		}

		if m.InstanceGroups[i] != nil {

			if err := m.InstanceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StackV2Request) validateNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {

		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *StackV2Request) validatePlacement(formats strfmt.Registry) error {

	if swag.IsZero(m.Placement) { // not required
		return nil
	}

	if m.Placement != nil {

		if err := m.Placement.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("placement")
			}
			return err
		}
	}

	return nil
}

func (m *StackV2Request) validateStackAuthentication(formats strfmt.Registry) error {

	if swag.IsZero(m.StackAuthentication) { // not required
		return nil
	}

	if m.StackAuthentication != nil {

		if err := m.StackAuthentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stackAuthentication")
			}
			return err
		}
	}

	return nil
}

func (m *StackV2Request) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if m.Tags != nil {

		if err := m.Tags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackV2Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackV2Request) UnmarshalBinary(b []byte) error {
	var res StackV2Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
