package templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models"
)

// NewPostTemplatesUserParams creates a new PostTemplatesUserParams object
// with the default values initialized.
func NewPostTemplatesUserParams() *PostTemplatesUserParams {
	var ()
	return &PostTemplatesUserParams{}
}

/*PostTemplatesUserParams contains all the parameters to send to the API endpoint
for the post templates user operation typically these are written to a http.Request
*/
type PostTemplatesUserParams struct {

	/*Body*/
	Body *models.TemplateRequest
}

// WithBody adds the body to the post templates user params
func (o *PostTemplatesUserParams) WithBody(body *models.TemplateRequest) *PostTemplatesUserParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostTemplatesUserParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.TemplateRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
