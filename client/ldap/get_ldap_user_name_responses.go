package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models"
)

// GetLdapUserNameReader is a Reader for the GetLdapUserName structure.
type GetLdapUserNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetLdapUserNameReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLdapUserNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLdapUserNameOK creates a GetLdapUserNameOK with default headers values
func NewGetLdapUserNameOK() *GetLdapUserNameOK {
	return &GetLdapUserNameOK{}
}

/*GetLdapUserNameOK handles this case with default header values.

successful operation
*/
type GetLdapUserNameOK struct {
	Payload *models.LdapConfigResponse
}

func (o *GetLdapUserNameOK) Error() string {
	return fmt.Sprintf("[GET /ldap/user/{name}][%d] getLdapUserNameOK  %+v", 200, o.Payload)
}

func (o *GetLdapUserNameOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LdapConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
