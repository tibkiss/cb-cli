package constraints

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

// GetConstraintsAccountNameReader is a Reader for the GetConstraintsAccountName structure.
type GetConstraintsAccountNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetConstraintsAccountNameReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetConstraintsAccountNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetConstraintsAccountNameOK creates a GetConstraintsAccountNameOK with default headers values
func NewGetConstraintsAccountNameOK() *GetConstraintsAccountNameOK {
	return &GetConstraintsAccountNameOK{}
}

/*GetConstraintsAccountNameOK handles this case with default header values.

successful operation
*/
type GetConstraintsAccountNameOK struct {
	Payload *models.ConstraintTemplateResponse
}

func (o *GetConstraintsAccountNameOK) Error() string {
	return fmt.Sprintf("[GET /constraints/account/{name}][%d] getConstraintsAccountNameOK  %+v", 200, o.Payload)
}

func (o *GetConstraintsAccountNameOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstraintTemplateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
