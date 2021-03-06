// Code generated by go-swagger; DO NOT EDIT.

package v1clustertemplates

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

// NewPostPublicClusterTemplateParams creates a new PostPublicClusterTemplateParams object
// with the default values initialized.
func NewPostPublicClusterTemplateParams() *PostPublicClusterTemplateParams {
	var ()
	return &PostPublicClusterTemplateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPublicClusterTemplateParamsWithTimeout creates a new PostPublicClusterTemplateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPublicClusterTemplateParamsWithTimeout(timeout time.Duration) *PostPublicClusterTemplateParams {
	var ()
	return &PostPublicClusterTemplateParams{

		timeout: timeout,
	}
}

// NewPostPublicClusterTemplateParamsWithContext creates a new PostPublicClusterTemplateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPublicClusterTemplateParamsWithContext(ctx context.Context) *PostPublicClusterTemplateParams {
	var ()
	return &PostPublicClusterTemplateParams{

		Context: ctx,
	}
}

// NewPostPublicClusterTemplateParamsWithHTTPClient creates a new PostPublicClusterTemplateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPublicClusterTemplateParamsWithHTTPClient(client *http.Client) *PostPublicClusterTemplateParams {
	var ()
	return &PostPublicClusterTemplateParams{
		HTTPClient: client,
	}
}

/*PostPublicClusterTemplateParams contains all the parameters to send to the API endpoint
for the post public cluster template operation typically these are written to a http.Request
*/
type PostPublicClusterTemplateParams struct {

	/*Body*/
	Body *models_cloudbreak.ClusterTemplateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post public cluster template params
func (o *PostPublicClusterTemplateParams) WithTimeout(timeout time.Duration) *PostPublicClusterTemplateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post public cluster template params
func (o *PostPublicClusterTemplateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post public cluster template params
func (o *PostPublicClusterTemplateParams) WithContext(ctx context.Context) *PostPublicClusterTemplateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post public cluster template params
func (o *PostPublicClusterTemplateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post public cluster template params
func (o *PostPublicClusterTemplateParams) WithHTTPClient(client *http.Client) *PostPublicClusterTemplateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post public cluster template params
func (o *PostPublicClusterTemplateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post public cluster template params
func (o *PostPublicClusterTemplateParams) WithBody(body *models_cloudbreak.ClusterTemplateRequest) *PostPublicClusterTemplateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post public cluster template params
func (o *PostPublicClusterTemplateParams) SetBody(body *models_cloudbreak.ClusterTemplateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPublicClusterTemplateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.ClusterTemplateRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
