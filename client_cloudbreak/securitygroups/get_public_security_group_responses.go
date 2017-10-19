package securitygroups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPublicSecurityGroupReader is a Reader for the GetPublicSecurityGroup structure.
type GetPublicSecurityGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPublicSecurityGroupReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicSecurityGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicSecurityGroupOK creates a GetPublicSecurityGroupOK with default headers values
func NewGetPublicSecurityGroupOK() *GetPublicSecurityGroupOK {
	return &GetPublicSecurityGroupOK{}
}

/*GetPublicSecurityGroupOK handles this case with default header values.

successful operation
*/
type GetPublicSecurityGroupOK struct {
	Payload *models_cloudbreak.SecurityGroupResponse
}

func (o *GetPublicSecurityGroupOK) Error() string {
	return fmt.Sprintf("[GET /securitygroups/account/{name}][%d] getPublicSecurityGroupOK  %+v", 200, o.Payload)
}

func (o *GetPublicSecurityGroupOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.SecurityGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
