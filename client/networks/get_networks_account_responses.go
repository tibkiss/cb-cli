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

// GetNetworksAccountReader is a Reader for the GetNetworksAccount structure.
type GetNetworksAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetNetworksAccountReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetworksAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNetworksAccountOK creates a GetNetworksAccountOK with default headers values
func NewGetNetworksAccountOK() *GetNetworksAccountOK {
	return &GetNetworksAccountOK{}
}

/*GetNetworksAccountOK handles this case with default header values.

successful operation
*/
type GetNetworksAccountOK struct {
	Payload []*models.NetworkJSON
}

func (o *GetNetworksAccountOK) Error() string {
	return fmt.Sprintf("[GET /networks/account][%d] getNetworksAccountOK  %+v", 200, o.Payload)
}

func (o *GetNetworksAccountOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
