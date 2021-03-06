// Code generated by go-swagger; DO NOT EDIT.

package v1util

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

// NewTestRdsConnectionUtilParams creates a new TestRdsConnectionUtilParams object
// with the default values initialized.
func NewTestRdsConnectionUtilParams() *TestRdsConnectionUtilParams {
	var ()
	return &TestRdsConnectionUtilParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestRdsConnectionUtilParamsWithTimeout creates a new TestRdsConnectionUtilParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestRdsConnectionUtilParamsWithTimeout(timeout time.Duration) *TestRdsConnectionUtilParams {
	var ()
	return &TestRdsConnectionUtilParams{

		timeout: timeout,
	}
}

// NewTestRdsConnectionUtilParamsWithContext creates a new TestRdsConnectionUtilParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestRdsConnectionUtilParamsWithContext(ctx context.Context) *TestRdsConnectionUtilParams {
	var ()
	return &TestRdsConnectionUtilParams{

		Context: ctx,
	}
}

// NewTestRdsConnectionUtilParamsWithHTTPClient creates a new TestRdsConnectionUtilParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestRdsConnectionUtilParamsWithHTTPClient(client *http.Client) *TestRdsConnectionUtilParams {
	var ()
	return &TestRdsConnectionUtilParams{
		HTTPClient: client,
	}
}

/*TestRdsConnectionUtilParams contains all the parameters to send to the API endpoint
for the test rds connection util operation typically these are written to a http.Request
*/
type TestRdsConnectionUtilParams struct {

	/*Body*/
	Body *models_cloudbreak.RdsConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test rds connection util params
func (o *TestRdsConnectionUtilParams) WithTimeout(timeout time.Duration) *TestRdsConnectionUtilParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test rds connection util params
func (o *TestRdsConnectionUtilParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test rds connection util params
func (o *TestRdsConnectionUtilParams) WithContext(ctx context.Context) *TestRdsConnectionUtilParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test rds connection util params
func (o *TestRdsConnectionUtilParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test rds connection util params
func (o *TestRdsConnectionUtilParams) WithHTTPClient(client *http.Client) *TestRdsConnectionUtilParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test rds connection util params
func (o *TestRdsConnectionUtilParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the test rds connection util params
func (o *TestRdsConnectionUtilParams) WithBody(body *models_cloudbreak.RdsConfig) *TestRdsConnectionUtilParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the test rds connection util params
func (o *TestRdsConnectionUtilParams) SetBody(body *models_cloudbreak.RdsConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TestRdsConnectionUtilParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
