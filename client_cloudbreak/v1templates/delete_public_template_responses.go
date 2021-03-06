// Code generated by go-swagger; DO NOT EDIT.

package v1templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePublicTemplateReader is a Reader for the DeletePublicTemplate structure.
type DeletePublicTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePublicTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeletePublicTemplateDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeletePublicTemplateDefault creates a DeletePublicTemplateDefault with default headers values
func NewDeletePublicTemplateDefault(code int) *DeletePublicTemplateDefault {
	return &DeletePublicTemplateDefault{
		_statusCode: code,
	}
}

/*DeletePublicTemplateDefault handles this case with default header values.

successful operation
*/
type DeletePublicTemplateDefault struct {
	_statusCode int
}

// Code gets the status code for the delete public template default response
func (o *DeletePublicTemplateDefault) Code() int {
	return o._statusCode
}

func (o *DeletePublicTemplateDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/templates/account/{name}][%d] deletePublicTemplate default ", o._statusCode)
}

func (o *DeletePublicTemplateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
