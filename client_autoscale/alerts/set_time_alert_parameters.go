package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/cb-cli/models_autoscale"
)

// NewSetTimeAlertParams creates a new SetTimeAlertParams object
// with the default values initialized.
func NewSetTimeAlertParams() *SetTimeAlertParams {
	var ()
	return &SetTimeAlertParams{}
}

/*SetTimeAlertParams contains all the parameters to send to the API endpoint
for the set time alert operation typically these are written to a http.Request
*/
type SetTimeAlertParams struct {

	/*AlertID*/
	AlertID int64
	/*Body*/
	Body *models_autoscale.MetricSchedule
	/*ClusterID*/
	ClusterID int64
}

// WithAlertID adds the alertId to the set time alert params
func (o *SetTimeAlertParams) WithAlertID(alertId int64) *SetTimeAlertParams {
	o.AlertID = alertId
	return o
}

// WithBody adds the body to the set time alert params
func (o *SetTimeAlertParams) WithBody(body *models_autoscale.MetricSchedule) *SetTimeAlertParams {
	o.Body = body
	return o
}

// WithClusterID adds the clusterId to the set time alert params
func (o *SetTimeAlertParams) WithClusterID(clusterId int64) *SetTimeAlertParams {
	o.ClusterID = clusterId
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *SetTimeAlertParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param alertId
	if err := r.SetPathParam("alertId", swag.FormatInt64(o.AlertID)); err != nil {
		return err
	}

	if o.Body == nil {
		o.Body = new(models_autoscale.MetricSchedule)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param clusterId
	if err := r.SetPathParam("clusterId", swag.FormatInt64(o.ClusterID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
