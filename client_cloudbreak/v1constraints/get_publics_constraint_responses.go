// Code generated by go-swagger; DO NOT EDIT.

package v1constraints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPublicsConstraintReader is a Reader for the GetPublicsConstraint structure.
type GetPublicsConstraintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicsConstraintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicsConstraintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicsConstraintOK creates a GetPublicsConstraintOK with default headers values
func NewGetPublicsConstraintOK() *GetPublicsConstraintOK {
	return &GetPublicsConstraintOK{}
}

/*GetPublicsConstraintOK handles this case with default header values.

successful operation
*/
type GetPublicsConstraintOK struct {
	Payload []*models_cloudbreak.ConstraintTemplateResponse
}

func (o *GetPublicsConstraintOK) Error() string {
	return fmt.Sprintf("[GET /v1/constraints/account][%d] getPublicsConstraintOK  %+v", 200, o.Payload)
}

func (o *GetPublicsConstraintOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
