// Code generated by go-swagger; DO NOT EDIT.

package v1users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ModifyProfileReader is a Reader for the ModifyProfile structure.
type ModifyProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewModifyProfileDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewModifyProfileDefault creates a ModifyProfileDefault with default headers values
func NewModifyProfileDefault(code int) *ModifyProfileDefault {
	return &ModifyProfileDefault{
		_statusCode: code,
	}
}

/*ModifyProfileDefault handles this case with default header values.

successful operation
*/
type ModifyProfileDefault struct {
	_statusCode int
}

// Code gets the status code for the modify profile default response
func (o *ModifyProfileDefault) Code() int {
	return o._statusCode
}

func (o *ModifyProfileDefault) Error() string {
	return fmt.Sprintf("[PUT /v1/users/profile][%d] modifyProfile default ", o._statusCode)
}

func (o *ModifyProfileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
