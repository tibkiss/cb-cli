package rdsconfigs

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

// GetRdsconfigsUserNameReader is a Reader for the GetRdsconfigsUserName structure.
type GetRdsconfigsUserNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetRdsconfigsUserNameReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRdsconfigsUserNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRdsconfigsUserNameOK creates a GetRdsconfigsUserNameOK with default headers values
func NewGetRdsconfigsUserNameOK() *GetRdsconfigsUserNameOK {
	return &GetRdsconfigsUserNameOK{}
}

/*GetRdsconfigsUserNameOK handles this case with default header values.

successful operation
*/
type GetRdsconfigsUserNameOK struct {
	Payload *models.RDSConfig
}

func (o *GetRdsconfigsUserNameOK) Error() string {
	return fmt.Sprintf("[GET /rdsconfigs/user/{name}][%d] getRdsconfigsUserNameOK  %+v", 200, o.Payload)
}

func (o *GetRdsconfigsUserNameOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RDSConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
