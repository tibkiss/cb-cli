package events

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

// GetEventsReader is a Reader for the GetEvents structure.
type GetEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetEventsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEventsOK creates a GetEventsOK with default headers values
func NewGetEventsOK() *GetEventsOK {
	return &GetEventsOK{}
}

/*GetEventsOK handles this case with default header values.

successful operation
*/
type GetEventsOK struct {
	Payload []*models_cloudbreak.CloudbreakEvent
}

func (o *GetEventsOK) Error() string {
	return fmt.Sprintf("[GET /events][%d] getEventsOK  %+v", 200, o.Payload)
}

func (o *GetEventsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
