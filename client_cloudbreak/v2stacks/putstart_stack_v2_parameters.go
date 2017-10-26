// Code generated by go-swagger; DO NOT EDIT.

package v2stacks

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
)

// NewPutstartStackV2Params creates a new PutstartStackV2Params object
// with the default values initialized.
func NewPutstartStackV2Params() *PutstartStackV2Params {
	var ()
	return &PutstartStackV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutstartStackV2ParamsWithTimeout creates a new PutstartStackV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutstartStackV2ParamsWithTimeout(timeout time.Duration) *PutstartStackV2Params {
	var ()
	return &PutstartStackV2Params{

		timeout: timeout,
	}
}

// NewPutstartStackV2ParamsWithContext creates a new PutstartStackV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewPutstartStackV2ParamsWithContext(ctx context.Context) *PutstartStackV2Params {
	var ()
	return &PutstartStackV2Params{

		Context: ctx,
	}
}

// NewPutstartStackV2ParamsWithHTTPClient creates a new PutstartStackV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutstartStackV2ParamsWithHTTPClient(client *http.Client) *PutstartStackV2Params {
	var ()
	return &PutstartStackV2Params{
		HTTPClient: client,
	}
}

/*PutstartStackV2Params contains all the parameters to send to the API endpoint
for the putstart stack v2 operation typically these are written to a http.Request
*/
type PutstartStackV2Params struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the putstart stack v2 params
func (o *PutstartStackV2Params) WithTimeout(timeout time.Duration) *PutstartStackV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the putstart stack v2 params
func (o *PutstartStackV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the putstart stack v2 params
func (o *PutstartStackV2Params) WithContext(ctx context.Context) *PutstartStackV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the putstart stack v2 params
func (o *PutstartStackV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the putstart stack v2 params
func (o *PutstartStackV2Params) WithHTTPClient(client *http.Client) *PutstartStackV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the putstart stack v2 params
func (o *PutstartStackV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the putstart stack v2 params
func (o *PutstartStackV2Params) WithName(name string) *PutstartStackV2Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the putstart stack v2 params
func (o *PutstartStackV2Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PutstartStackV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
