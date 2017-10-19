package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewPostPublicCredentialParams creates a new PostPublicCredentialParams object
// with the default values initialized.
func NewPostPublicCredentialParams() *PostPublicCredentialParams {
	var ()
	return &PostPublicCredentialParams{}
}

/*PostPublicCredentialParams contains all the parameters to send to the API endpoint
for the post public credential operation typically these are written to a http.Request
*/
type PostPublicCredentialParams struct {

	/*Body*/
	Body *models_cloudbreak.CredentialRequest
}

// WithBody adds the body to the post public credential params
func (o *PostPublicCredentialParams) WithBody(body *models_cloudbreak.CredentialRequest) *PostPublicCredentialParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostPublicCredentialParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.CredentialRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
