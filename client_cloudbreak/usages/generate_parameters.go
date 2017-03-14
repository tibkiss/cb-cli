package usages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGenerateParams creates a new GenerateParams object
// with the default values initialized.
func NewGenerateParams() *GenerateParams {

	return &GenerateParams{}
}

/*GenerateParams contains all the parameters to send to the API endpoint
for the generate operation typically these are written to a http.Request
*/
type GenerateParams struct {
}

// WriteToRequest writes these params to a swagger request
func (o *GenerateParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}