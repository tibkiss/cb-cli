package flexsubscriptions

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

// GetFlexSubscriptionByIDReader is a Reader for the GetFlexSubscriptionByID structure.
type GetFlexSubscriptionByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetFlexSubscriptionByIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFlexSubscriptionByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFlexSubscriptionByIDOK creates a GetFlexSubscriptionByIDOK with default headers values
func NewGetFlexSubscriptionByIDOK() *GetFlexSubscriptionByIDOK {
	return &GetFlexSubscriptionByIDOK{}
}

/*GetFlexSubscriptionByIDOK handles this case with default header values.

successful operation
*/
type GetFlexSubscriptionByIDOK struct {
	Payload *models_cloudbreak.FlexSubscriptionResponse
}

func (o *GetFlexSubscriptionByIDOK) Error() string {
	return fmt.Sprintf("[GET /flexsubscriptions/{id}][%d] getFlexSubscriptionByIdOK  %+v", 200, o.Payload)
}

func (o *GetFlexSubscriptionByIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.FlexSubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
