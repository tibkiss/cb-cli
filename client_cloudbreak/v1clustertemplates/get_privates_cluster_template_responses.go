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

// GetPrivatesClusterTemplateReader is a Reader for the GetPrivatesClusterTemplate structure.
type GetPrivatesClusterTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPrivatesClusterTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivatesClusterTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivatesClusterTemplateOK creates a GetPrivatesClusterTemplateOK with default headers values
func NewGetPrivatesClusterTemplateOK() *GetPrivatesClusterTemplateOK {
	return &GetPrivatesClusterTemplateOK{}
}

/*GetPrivatesClusterTemplateOK handles this case with default header values.

successful operation
*/
type GetPrivatesClusterTemplateOK struct {
	Payload []*models_cloudbreak.ClusterTemplateResponse
}

func (o *GetPrivatesClusterTemplateOK) Error() string {
	return fmt.Sprintf("[GET /v1/clustertemplates/user][%d] getPrivatesClusterTemplateOK  %+v", 200, o.Payload)
}

func (o *GetPrivatesClusterTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
