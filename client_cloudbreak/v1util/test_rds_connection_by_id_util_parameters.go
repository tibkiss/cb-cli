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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTestRdsConnectionByIDUtilParams creates a new TestRdsConnectionByIDUtilParams object
// with the default values initialized.
func NewTestRdsConnectionByIDUtilParams() *TestRdsConnectionByIDUtilParams {
	var ()
	return &TestRdsConnectionByIDUtilParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestRdsConnectionByIDUtilParamsWithTimeout creates a new TestRdsConnectionByIDUtilParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestRdsConnectionByIDUtilParamsWithTimeout(timeout time.Duration) *TestRdsConnectionByIDUtilParams {
	var ()
	return &TestRdsConnectionByIDUtilParams{

		timeout: timeout,
	}
}

// NewTestRdsConnectionByIDUtilParamsWithContext creates a new TestRdsConnectionByIDUtilParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestRdsConnectionByIDUtilParamsWithContext(ctx context.Context) *TestRdsConnectionByIDUtilParams {
	var ()
	return &TestRdsConnectionByIDUtilParams{

		Context: ctx,
	}
}

// NewTestRdsConnectionByIDUtilParamsWithHTTPClient creates a new TestRdsConnectionByIDUtilParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestRdsConnectionByIDUtilParamsWithHTTPClient(client *http.Client) *TestRdsConnectionByIDUtilParams {
	var ()
	return &TestRdsConnectionByIDUtilParams{
		HTTPClient: client,
	}
}

/*TestRdsConnectionByIDUtilParams contains all the parameters to send to the API endpoint
for the test rds connection by Id util operation typically these are written to a http.Request
*/
type TestRdsConnectionByIDUtilParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test rds connection by Id util params
func (o *TestRdsConnectionByIDUtilParams) WithTimeout(timeout time.Duration) *TestRdsConnectionByIDUtilParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test rds connection by Id util params
func (o *TestRdsConnectionByIDUtilParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test rds connection by Id util params
func (o *TestRdsConnectionByIDUtilParams) WithContext(ctx context.Context) *TestRdsConnectionByIDUtilParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test rds connection by Id util params
func (o *TestRdsConnectionByIDUtilParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test rds connection by Id util params
func (o *TestRdsConnectionByIDUtilParams) WithHTTPClient(client *http.Client) *TestRdsConnectionByIDUtilParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test rds connection by Id util params
func (o *TestRdsConnectionByIDUtilParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the test rds connection by Id util params
func (o *TestRdsConnectionByIDUtilParams) WithID(id int64) *TestRdsConnectionByIDUtilParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the test rds connection by Id util params
func (o *TestRdsConnectionByIDUtilParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TestRdsConnectionByIDUtilParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
