// Code generated by go-swagger; DO NOT EDIT.

package v2stacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// StatusStackV2Reader is a Reader for the StatusStackV2 structure.
type StatusStackV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StatusStackV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStatusStackV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewStatusStackV2OK creates a StatusStackV2OK with default headers values
func NewStatusStackV2OK() *StatusStackV2OK {
	return &StatusStackV2OK{}
}

/*StatusStackV2OK handles this case with default header values.

successful operation
*/
type StatusStackV2OK struct {
	Payload StatusStackV2OKBody
}

func (o *StatusStackV2OK) Error() string {
	return fmt.Sprintf("[GET /v2/stacks/{id}/status][%d] statusStackV2OK  %+v", 200, o.Payload)
}

func (o *StatusStackV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StatusStackV2OKBody status stack v2 o k body
swagger:model StatusStackV2OKBody
*/

type StatusStackV2OKBody map[string]interface{}

// Validate validates this status stack v2 o k body
func (o StatusStackV2OKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Required("statusStackV2OK", "body", o); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
