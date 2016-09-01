package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models"
)

// NewPostLdapUserParams creates a new PostLdapUserParams object
// with the default values initialized.
func NewPostLdapUserParams() *PostLdapUserParams {
	var ()
	return &PostLdapUserParams{}
}

/*PostLdapUserParams contains all the parameters to send to the API endpoint
for the post ldap user operation typically these are written to a http.Request
*/
type PostLdapUserParams struct {

	/*Body*/
	Body *models.LdapConfigRequest
}

// WithBody adds the body to the post ldap user params
func (o *PostLdapUserParams) WithBody(body *models.LdapConfigRequest) *PostLdapUserParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostLdapUserParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.LdapConfigRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
