package topologies

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

// PostTopologiesAccountReader is a Reader for the PostTopologiesAccount structure.
type PostTopologiesAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostTopologiesAccountReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostTopologiesAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTopologiesAccountOK creates a PostTopologiesAccountOK with default headers values
func NewPostTopologiesAccountOK() *PostTopologiesAccountOK {
	return &PostTopologiesAccountOK{}
}

/*PostTopologiesAccountOK handles this case with default header values.

successful operation
*/
type PostTopologiesAccountOK struct {
	Payload *models.ID
}

func (o *PostTopologiesAccountOK) Error() string {
	return fmt.Sprintf("[POST /topologies/account][%d] postTopologiesAccountOK  %+v", 200, o.Payload)
}

func (o *PostTopologiesAccountOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
