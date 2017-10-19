package constraints

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

// PostPrivateConstraintReader is a Reader for the PostPrivateConstraint structure.
type PostPrivateConstraintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostPrivateConstraintReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostPrivateConstraintOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostPrivateConstraintOK creates a PostPrivateConstraintOK with default headers values
func NewPostPrivateConstraintOK() *PostPrivateConstraintOK {
	return &PostPrivateConstraintOK{}
}

/*PostPrivateConstraintOK handles this case with default header values.

successful operation
*/
type PostPrivateConstraintOK struct {
	Payload *models_cloudbreak.ConstraintTemplateResponse
}

func (o *PostPrivateConstraintOK) Error() string {
	return fmt.Sprintf("[POST /constraints/user][%d] postPrivateConstraintOK  %+v", 200, o.Payload)
}

func (o *PostPrivateConstraintOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.ConstraintTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
