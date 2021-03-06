// Code generated by go-swagger; DO NOT EDIT.

package v1smartsensesubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// PostPrivateSmartSenseSubscriptionReader is a Reader for the PostPrivateSmartSenseSubscription structure.
type PostPrivateSmartSenseSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostPrivateSmartSenseSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPrivateSmartSenseSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPrivateSmartSenseSubscriptionOK creates a PostPrivateSmartSenseSubscriptionOK with default headers values
func NewPostPrivateSmartSenseSubscriptionOK() *PostPrivateSmartSenseSubscriptionOK {
	return &PostPrivateSmartSenseSubscriptionOK{}
}

/*PostPrivateSmartSenseSubscriptionOK handles this case with default header values.

successful operation
*/
type PostPrivateSmartSenseSubscriptionOK struct {
	Payload *models_cloudbreak.SmartSenseSubscriptionJSON
}

func (o *PostPrivateSmartSenseSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /v1/smartsensesubscriptions/user][%d] postPrivateSmartSenseSubscriptionOK  %+v", 200, o.Payload)
}

func (o *PostPrivateSmartSenseSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.SmartSenseSubscriptionJSON)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
