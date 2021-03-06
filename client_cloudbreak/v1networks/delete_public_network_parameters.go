// Code generated by go-swagger; DO NOT EDIT.

package v1networks

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

// NewDeletePublicNetworkParams creates a new DeletePublicNetworkParams object
// with the default values initialized.
func NewDeletePublicNetworkParams() *DeletePublicNetworkParams {
	var ()
	return &DeletePublicNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePublicNetworkParamsWithTimeout creates a new DeletePublicNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePublicNetworkParamsWithTimeout(timeout time.Duration) *DeletePublicNetworkParams {
	var ()
	return &DeletePublicNetworkParams{

		timeout: timeout,
	}
}

// NewDeletePublicNetworkParamsWithContext creates a new DeletePublicNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePublicNetworkParamsWithContext(ctx context.Context) *DeletePublicNetworkParams {
	var ()
	return &DeletePublicNetworkParams{

		Context: ctx,
	}
}

// NewDeletePublicNetworkParamsWithHTTPClient creates a new DeletePublicNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePublicNetworkParamsWithHTTPClient(client *http.Client) *DeletePublicNetworkParams {
	var ()
	return &DeletePublicNetworkParams{
		HTTPClient: client,
	}
}

/*DeletePublicNetworkParams contains all the parameters to send to the API endpoint
for the delete public network operation typically these are written to a http.Request
*/
type DeletePublicNetworkParams struct {

	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete public network params
func (o *DeletePublicNetworkParams) WithTimeout(timeout time.Duration) *DeletePublicNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete public network params
func (o *DeletePublicNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete public network params
func (o *DeletePublicNetworkParams) WithContext(ctx context.Context) *DeletePublicNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete public network params
func (o *DeletePublicNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete public network params
func (o *DeletePublicNetworkParams) WithHTTPClient(client *http.Client) *DeletePublicNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete public network params
func (o *DeletePublicNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete public network params
func (o *DeletePublicNetworkParams) WithName(name string) *DeletePublicNetworkParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete public network params
func (o *DeletePublicNetworkParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePublicNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
