// Code generated by go-swagger; DO NOT EDIT.

package v2stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteInstanceStackV2Reader is a Reader for the DeleteInstanceStackV2 structure.
type DeleteInstanceStackV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInstanceStackV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeleteInstanceStackV2Default(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeleteInstanceStackV2Default creates a DeleteInstanceStackV2Default with default headers values
func NewDeleteInstanceStackV2Default(code int) *DeleteInstanceStackV2Default {
	return &DeleteInstanceStackV2Default{
		_statusCode: code,
	}
}

/*DeleteInstanceStackV2Default handles this case with default header values.

successful operation
*/
type DeleteInstanceStackV2Default struct {
	_statusCode int
}

// Code gets the status code for the delete instance stack v2 default response
func (o *DeleteInstanceStackV2Default) Code() int {
	return o._statusCode
}

func (o *DeleteInstanceStackV2Default) Error() string {
	return fmt.Sprintf("[DELETE /v2/stacks/{stackId}/{instanceId}][%d] deleteInstanceStackV2 default ", o._statusCode)
}

func (o *DeleteInstanceStackV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
