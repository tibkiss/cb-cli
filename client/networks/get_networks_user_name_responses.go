package networks

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

// GetNetworksUserNameReader is a Reader for the GetNetworksUserName structure.
type GetNetworksUserNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetNetworksUserNameReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetworksUserNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworksUserNameOK creates a GetNetworksUserNameOK with default headers values
func NewGetNetworksUserNameOK() *GetNetworksUserNameOK {
	return &GetNetworksUserNameOK{}
}

/*GetNetworksUserNameOK handles this case with default header values.

successful operation
*/
type GetNetworksUserNameOK struct {
	Payload *models.NetworkJSON
}

func (o *GetNetworksUserNameOK) Error() string {
	return fmt.Sprintf("[GET /networks/user/{name}][%d] getNetworksUserNameOK  %+v", 200, o.Payload)
}

func (o *GetNetworksUserNameOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
