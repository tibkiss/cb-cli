// Code generated by go-swagger; DO NOT EDIT.

package v1stacks

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

// NewDeleteStackParams creates a new DeleteStackParams object
// with the default values initialized.
func NewDeleteStackParams() *DeleteStackParams {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeleteStackParams{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteStackParamsWithTimeout creates a new DeleteStackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteStackParamsWithTimeout(timeout time.Duration) *DeleteStackParams {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeleteStackParams{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,

		timeout: timeout,
	}
}

// NewDeleteStackParamsWithContext creates a new DeleteStackParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteStackParamsWithContext(ctx context.Context) *DeleteStackParams {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeleteStackParams{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,

		Context: ctx,
	}
}

// NewDeleteStackParamsWithHTTPClient creates a new DeleteStackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteStackParamsWithHTTPClient(client *http.Client) *DeleteStackParams {
	var (
		deleteDependenciesDefault = bool(false)
		forcedDefault             = bool(false)
	)
	return &DeleteStackParams{
		DeleteDependencies: &deleteDependenciesDefault,
		Forced:             &forcedDefault,
		HTTPClient:         client,
	}
}

/*DeleteStackParams contains all the parameters to send to the API endpoint
for the delete stack operation typically these are written to a http.Request
*/
type DeleteStackParams struct {

	/*DeleteDependencies*/
	DeleteDependencies *bool
	/*Forced*/
	Forced *bool
	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete stack params
func (o *DeleteStackParams) WithTimeout(timeout time.Duration) *DeleteStackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete stack params
func (o *DeleteStackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete stack params
func (o *DeleteStackParams) WithContext(ctx context.Context) *DeleteStackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete stack params
func (o *DeleteStackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete stack params
func (o *DeleteStackParams) WithHTTPClient(client *http.Client) *DeleteStackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete stack params
func (o *DeleteStackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleteDependencies adds the deleteDependencies to the delete stack params
func (o *DeleteStackParams) WithDeleteDependencies(deleteDependencies *bool) *DeleteStackParams {
	o.SetDeleteDependencies(deleteDependencies)
	return o
}

// SetDeleteDependencies adds the deleteDependencies to the delete stack params
func (o *DeleteStackParams) SetDeleteDependencies(deleteDependencies *bool) {
	o.DeleteDependencies = deleteDependencies
}

// WithForced adds the forced to the delete stack params
func (o *DeleteStackParams) WithForced(forced *bool) *DeleteStackParams {
	o.SetForced(forced)
	return o
}

// SetForced adds the forced to the delete stack params
func (o *DeleteStackParams) SetForced(forced *bool) {
	o.Forced = forced
}

// WithID adds the id to the delete stack params
func (o *DeleteStackParams) WithID(id int64) *DeleteStackParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete stack params
func (o *DeleteStackParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteStackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DeleteDependencies != nil {

		// query param deleteDependencies
		var qrDeleteDependencies bool
		if o.DeleteDependencies != nil {
			qrDeleteDependencies = *o.DeleteDependencies
		}
		qDeleteDependencies := swag.FormatBool(qrDeleteDependencies)
		if qDeleteDependencies != "" {
			if err := r.SetQueryParam("deleteDependencies", qDeleteDependencies); err != nil {
				return err
			}
		}

	}

	if o.Forced != nil {

		// query param forced
		var qrForced bool
		if o.Forced != nil {
			qrForced = *o.Forced
		}
		qForced := swag.FormatBool(qrForced)
		if qForced != "" {
			if err := r.SetQueryParam("forced", qForced); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
