// Code generated by go-swagger; DO NOT EDIT.

package v1clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RepairClusterReader is a Reader for the RepairCluster structure.
type RepairClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RepairClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewRepairClusterDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewRepairClusterDefault creates a RepairClusterDefault with default headers values
func NewRepairClusterDefault(code int) *RepairClusterDefault {
	return &RepairClusterDefault{
		_statusCode: code,
	}
}

/*RepairClusterDefault handles this case with default header values.

successful operation
*/
type RepairClusterDefault struct {
	_statusCode int
}

// Code gets the status code for the repair cluster default response
func (o *RepairClusterDefault) Code() int {
	return o._statusCode
}

func (o *RepairClusterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{id}/cluster/manualrepair][%d] repairCluster default ", o._statusCode)
}

func (o *RepairClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}