package blueprints

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

// GetBlueprintsIDReader is a Reader for the GetBlueprintsID structure.
type GetBlueprintsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetBlueprintsIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBlueprintsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBlueprintsIDOK creates a GetBlueprintsIDOK with default headers values
func NewGetBlueprintsIDOK() *GetBlueprintsIDOK {
	return &GetBlueprintsIDOK{}
}

/*GetBlueprintsIDOK handles this case with default header values.

successful operation
*/
type GetBlueprintsIDOK struct {
	Payload *models.BlueprintResponse
}

func (o *GetBlueprintsIDOK) Error() string {
	return fmt.Sprintf("[GET /blueprints/{id}][%d] getBlueprintsIdOK  %+v", 200, o.Payload)
}

func (o *GetBlueprintsIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BlueprintResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
