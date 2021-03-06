// Code generated by go-swagger; DO NOT EDIT.

package v1rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewPostPrivateRdsParams creates a new PostPrivateRdsParams object
// with the default values initialized.
func NewPostPrivateRdsParams() *PostPrivateRdsParams {
	var ()
	return &PostPrivateRdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPrivateRdsParamsWithTimeout creates a new PostPrivateRdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPrivateRdsParamsWithTimeout(timeout time.Duration) *PostPrivateRdsParams {
	var ()
	return &PostPrivateRdsParams{

		timeout: timeout,
	}
}

// NewPostPrivateRdsParamsWithContext creates a new PostPrivateRdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPrivateRdsParamsWithContext(ctx context.Context) *PostPrivateRdsParams {
	var ()
	return &PostPrivateRdsParams{

		Context: ctx,
	}
}

// NewPostPrivateRdsParamsWithHTTPClient creates a new PostPrivateRdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPrivateRdsParamsWithHTTPClient(client *http.Client) *PostPrivateRdsParams {
	var ()
	return &PostPrivateRdsParams{
		HTTPClient: client,
	}
}

/*PostPrivateRdsParams contains all the parameters to send to the API endpoint
for the post private rds operation typically these are written to a http.Request
*/
type PostPrivateRdsParams struct {

	/*Body*/
	Body *models_cloudbreak.RdsConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post private rds params
func (o *PostPrivateRdsParams) WithTimeout(timeout time.Duration) *PostPrivateRdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post private rds params
func (o *PostPrivateRdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post private rds params
func (o *PostPrivateRdsParams) WithContext(ctx context.Context) *PostPrivateRdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post private rds params
func (o *PostPrivateRdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post private rds params
func (o *PostPrivateRdsParams) WithHTTPClient(client *http.Client) *PostPrivateRdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post private rds params
func (o *PostPrivateRdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post private rds params
func (o *PostPrivateRdsParams) WithBody(body *models_cloudbreak.RdsConfig) *PostPrivateRdsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post private rds params
func (o *PostPrivateRdsParams) SetBody(body *models_cloudbreak.RdsConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPrivateRdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.RdsConfig)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
