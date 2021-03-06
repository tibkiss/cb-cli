// Code generated by go-swagger; DO NOT EDIT.

package v1securitygroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteSecurityGroupReader is a Reader for the DeleteSecurityGroup structure.
type DeleteSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSecurityGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteSecurityGroupDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteSecurityGroupDefault creates a DeleteSecurityGroupDefault with default headers values
func NewDeleteSecurityGroupDefault(code int) *DeleteSecurityGroupDefault {
	return &DeleteSecurityGroupDefault{
		_statusCode: code,
	}
}

/*DeleteSecurityGroupDefault handles this case with default header values.

successful operation
*/
type DeleteSecurityGroupDefault struct {
	_statusCode int
}

// Code gets the status code for the delete security group default response
func (o *DeleteSecurityGroupDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSecurityGroupDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/securitygroups/{id}][%d] deleteSecurityGroup default ", o._statusCode)
}

func (o *DeleteSecurityGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
