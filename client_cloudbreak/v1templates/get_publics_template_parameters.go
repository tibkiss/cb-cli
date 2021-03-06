// Code generated by go-swagger; DO NOT EDIT.

package v1templates

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

// NewGetPublicsTemplateParams creates a new GetPublicsTemplateParams object
// with the default values initialized.
func NewGetPublicsTemplateParams() *GetPublicsTemplateParams {

	return &GetPublicsTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPublicsTemplateParamsWithTimeout creates a new GetPublicsTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPublicsTemplateParamsWithTimeout(timeout time.Duration) *GetPublicsTemplateParams {

	return &GetPublicsTemplateParams{

		timeout: timeout,
	}
}

// NewGetPublicsTemplateParamsWithContext creates a new GetPublicsTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPublicsTemplateParamsWithContext(ctx context.Context) *GetPublicsTemplateParams {

	return &GetPublicsTemplateParams{

		Context: ctx,
	}
}

// NewGetPublicsTemplateParamsWithHTTPClient creates a new GetPublicsTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPublicsTemplateParamsWithHTTPClient(client *http.Client) *GetPublicsTemplateParams {

	return &GetPublicsTemplateParams{
		HTTPClient: client,
	}
}

/*GetPublicsTemplateParams contains all the parameters to send to the API endpoint
for the get publics template operation typically these are written to a http.Request
*/
type GetPublicsTemplateParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get publics template params
func (o *GetPublicsTemplateParams) WithTimeout(timeout time.Duration) *GetPublicsTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get publics template params
func (o *GetPublicsTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get publics template params
func (o *GetPublicsTemplateParams) WithContext(ctx context.Context) *GetPublicsTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get publics template params
func (o *GetPublicsTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get publics template params
func (o *GetPublicsTemplateParams) WithHTTPClient(client *http.Client) *GetPublicsTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get publics template params
func (o *GetPublicsTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPublicsTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
