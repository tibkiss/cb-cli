package util

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models"
)

// NewTestAmbariDatabaseParams creates a new TestAmbariDatabaseParams object
// with the default values initialized.
func NewTestAmbariDatabaseParams() *TestAmbariDatabaseParams {
	var ()
	return &TestAmbariDatabaseParams{}
}

/*TestAmbariDatabaseParams contains all the parameters to send to the API endpoint
for the test ambari database operation typically these are written to a http.Request
*/
type TestAmbariDatabaseParams struct {

	/*Body*/
	Body *models.AmbariDatabaseDetails
}

// WithBody adds the body to the test ambari database params
func (o *TestAmbariDatabaseParams) WithBody(body *models.AmbariDatabaseDetails) *TestAmbariDatabaseParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *TestAmbariDatabaseParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.AmbariDatabaseDetails)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
