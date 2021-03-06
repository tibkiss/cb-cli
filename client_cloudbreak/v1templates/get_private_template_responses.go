// Code generated by go-swagger; DO NOT EDIT.

package v1templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPrivateTemplateReader is a Reader for the GetPrivateTemplate structure.
type GetPrivateTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivateTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateTemplateOK creates a GetPrivateTemplateOK with default headers values
func NewGetPrivateTemplateOK() *GetPrivateTemplateOK {
	return &GetPrivateTemplateOK{}
}

/*GetPrivateTemplateOK handles this case with default header values.

successful operation
*/
type GetPrivateTemplateOK struct {
	Payload *models_cloudbreak.TemplateResponse
}

func (o *GetPrivateTemplateOK) Error() string {
	return fmt.Sprintf("[GET /v1/templates/user/{name}][%d] getPrivateTemplateOK  %+v", 200, o.Payload)
}

func (o *GetPrivateTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.TemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
