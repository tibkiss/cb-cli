package stacks

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

// GetPublicsStackReader is a Reader for the GetPublicsStack structure.
type GetPublicsStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPublicsStackReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicsStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicsStackOK creates a GetPublicsStackOK with default headers values
func NewGetPublicsStackOK() *GetPublicsStackOK {
	return &GetPublicsStackOK{}
}

/*GetPublicsStackOK handles this case with default header values.

successful operation
*/
type GetPublicsStackOK struct {
	Payload []*models_cloudbreak.StackResponse
}

func (o *GetPublicsStackOK) Error() string {
	return fmt.Sprintf("[GET /stacks/account][%d] getPublicsStackOK  %+v", 200, o.Payload)
}

func (o *GetPublicsStackOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
