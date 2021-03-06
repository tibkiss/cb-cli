// Code generated by go-swagger; DO NOT EDIT.

package v1clustertemplates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPublicClusterTemplateReader is a Reader for the GetPublicClusterTemplate structure.
type GetPublicClusterTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicClusterTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicClusterTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicClusterTemplateOK creates a GetPublicClusterTemplateOK with default headers values
func NewGetPublicClusterTemplateOK() *GetPublicClusterTemplateOK {
	return &GetPublicClusterTemplateOK{}
}

/*GetPublicClusterTemplateOK handles this case with default header values.

successful operation
*/
type GetPublicClusterTemplateOK struct {
	Payload *models_cloudbreak.ClusterTemplateResponse
}

func (o *GetPublicClusterTemplateOK) Error() string {
	return fmt.Sprintf("[GET /v1/clustertemplates/account/{name}][%d] getPublicClusterTemplateOK  %+v", 200, o.Payload)
}

func (o *GetPublicClusterTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.ClusterTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
