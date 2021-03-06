// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RDSConfigResponse r d s config response
// swagger:model RDSConfigResponse

type RDSConfigResponse struct {

	// list of clusters which use config
	// Unique: true
	ClusterNames []string `json:"clusterNames"`

	// JDBC connection URL in the form of jdbc:<db-type>://<address>:<port>/<db>
	// Required: true
	// Pattern: ^jdbc:postgresql://[-\w\.]*:\d{1,5}/?\w*
	ConnectionURL *string `json:"connectionURL"`

	// creation time of the resource in long
	CreationDate int64 `json:"creationDate,omitempty"`

	// Type of the external database (allowed values: MYSQL, POSTGRES)
	// Required: true
	DatabaseType *string `json:"databaseType"`

	// HDP version for the RDS configuration
	// Required: true
	HdpVersion *string `json:"hdpVersion"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// Name of the RDS configuration resource
	// Required: true
	Name *string `json:"name"`

	// custom properties for rds connection
	// Unique: true
	Properties []*RdsConfigProperty `json:"properties"`

	// resource is visible in account
	PublicInAccount *bool `json:"publicInAccount,omitempty"`

	// Type of rds (HIVE or RANGER)
	Type string `json:"type,omitempty"`

	// If true, then the RDS configuration will be validated
	Validated *bool `json:"validated,omitempty"`
}

/* polymorph RDSConfigResponse clusterNames false */

/* polymorph RDSConfigResponse connectionURL false */

/* polymorph RDSConfigResponse creationDate false */

/* polymorph RDSConfigResponse databaseType false */

/* polymorph RDSConfigResponse hdpVersion false */

/* polymorph RDSConfigResponse id false */

/* polymorph RDSConfigResponse name false */

/* polymorph RDSConfigResponse properties false */

/* polymorph RDSConfigResponse publicInAccount false */

/* polymorph RDSConfigResponse type false */

/* polymorph RDSConfigResponse validated false */

// Validate validates this r d s config response
func (m *RDSConfigResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterNames(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectionURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDatabaseType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHdpVersion(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RDSConfigResponse) validateClusterNames(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterNames) { // not required
		return nil
	}

	if err := validate.UniqueItems("clusterNames", "body", m.ClusterNames); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateConnectionURL(formats strfmt.Registry) error {

	if err := validate.Required("connectionURL", "body", m.ConnectionURL); err != nil {
		return err
	}

	if err := validate.Pattern("connectionURL", "body", string(*m.ConnectionURL), `^jdbc:postgresql://[-\w\.]*:\d{1,5}/?\w*`); err != nil {
		return err
	}

	return nil
}

var rDSConfigResponseTypeDatabaseTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["POSTGRES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rDSConfigResponseTypeDatabaseTypePropEnum = append(rDSConfigResponseTypeDatabaseTypePropEnum, v)
	}
}

const (
	// RDSConfigResponseDatabaseTypePOSTGRES captures enum value "POSTGRES"
	RDSConfigResponseDatabaseTypePOSTGRES string = "POSTGRES"
)

// prop value enum
func (m *RDSConfigResponse) validateDatabaseTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, rDSConfigResponseTypeDatabaseTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RDSConfigResponse) validateDatabaseType(formats strfmt.Registry) error {

	if err := validate.Required("databaseType", "body", m.DatabaseType); err != nil {
		return err
	}

	// value enum
	if err := m.validateDatabaseTypeEnum("databaseType", "body", *m.DatabaseType); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateHdpVersion(formats strfmt.Registry) error {

	if err := validate.Required("hdpVersion", "body", m.HdpVersion); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *RDSConfigResponse) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if err := validate.UniqueItems("properties", "body", m.Properties); err != nil {
		return err
	}

	for i := 0; i < len(m.Properties); i++ {

		if swag.IsZero(m.Properties[i]) { // not required
			continue
		}

		if m.Properties[i] != nil {

			if err := m.Properties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("properties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var rDSConfigResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["HIVE","RANGER","DRUID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		rDSConfigResponseTypeTypePropEnum = append(rDSConfigResponseTypeTypePropEnum, v)
	}
}

const (
	// RDSConfigResponseTypeHIVE captures enum value "HIVE"
	RDSConfigResponseTypeHIVE string = "HIVE"
	// RDSConfigResponseTypeRANGER captures enum value "RANGER"
	RDSConfigResponseTypeRANGER string = "RANGER"
	// RDSConfigResponseTypeDRUID captures enum value "DRUID"
	RDSConfigResponseTypeDRUID string = "DRUID"
)

// prop value enum
func (m *RDSConfigResponse) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, rDSConfigResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RDSConfigResponse) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RDSConfigResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RDSConfigResponse) UnmarshalBinary(b []byte) error {
	var res RDSConfigResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
