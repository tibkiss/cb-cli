package topologies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewDeleteTopologiesAccountIDParams creates a new DeleteTopologiesAccountIDParams object
// with the default values initialized.
func NewDeleteTopologiesAccountIDParams() *DeleteTopologiesAccountIDParams {
	var (
		forcedDefault bool = bool(false)
	)
	return &DeleteTopologiesAccountIDParams{
		Forced: &forcedDefault,
	}
}

/*DeleteTopologiesAccountIDParams contains all the parameters to send to the API endpoint
for the delete topologies account ID operation typically these are written to a http.Request
*/
type DeleteTopologiesAccountIDParams struct {

	/*Forced*/
	Forced *bool
	/*ID*/
	ID int64
}

// WithForced adds the forced to the delete topologies account ID params
func (o *DeleteTopologiesAccountIDParams) WithForced(forced *bool) *DeleteTopologiesAccountIDParams {
	o.Forced = forced
	return o
}

// WithID adds the id to the delete topologies account ID params
func (o *DeleteTopologiesAccountIDParams) WithID(id int64) *DeleteTopologiesAccountIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTopologiesAccountIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

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