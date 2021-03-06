// Code generated by go-swagger; DO NOT EDIT.

package v1credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePublicCredentialReader is a Reader for the DeletePublicCredential structure.
type DeletePublicCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePublicCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeletePublicCredentialDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeletePublicCredentialDefault creates a DeletePublicCredentialDefault with default headers values
func NewDeletePublicCredentialDefault(code int) *DeletePublicCredentialDefault {
	return &DeletePublicCredentialDefault{
		_statusCode: code,
	}
}

/*DeletePublicCredentialDefault handles this case with default header values.

successful operation
*/
type DeletePublicCredentialDefault struct {
	_statusCode int
}

// Code gets the status code for the delete public credential default response
func (o *DeletePublicCredentialDefault) Code() int {
	return o._statusCode
}

func (o *DeletePublicCredentialDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/credentials/account/{name}][%d] deletePublicCredential default ", o._statusCode)
}

func (o *DeletePublicCredentialDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
