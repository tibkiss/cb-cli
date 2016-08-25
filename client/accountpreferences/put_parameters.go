package accountpreferences

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/sequenceiq/hdc-cli/models"
)

// NewPutParams creates a new PutParams object
// with the default values initialized.
func NewPutParams() *PutParams {
	var ()
	return &PutParams{}
}

/*PutParams contains all the parameters to send to the API endpoint
for the put operation typically these are written to a http.Request
*/
type PutParams struct {

	/*Body*/
	Body *models.AccountPreferences
}

// WithBody adds the body to the put params
func (o *PutParams) WithBody(body *models.AccountPreferences) *PutParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PutParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.AccountPreferences)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}