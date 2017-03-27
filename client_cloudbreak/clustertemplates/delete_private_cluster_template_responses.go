package clustertemplates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// DeletePrivateClusterTemplateReader is a Reader for the DeletePrivateClusterTemplate structure.
type DeletePrivateClusterTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeletePrivateClusterTemplateReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewDeletePrivateClusterTemplateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeletePrivateClusterTemplateDefault creates a DeletePrivateClusterTemplateDefault with default headers values
func NewDeletePrivateClusterTemplateDefault(code int) *DeletePrivateClusterTemplateDefault {
	return &DeletePrivateClusterTemplateDefault{
		_statusCode: code,
	}
}

/*DeletePrivateClusterTemplateDefault handles this case with default header values.

successful operation
*/
type DeletePrivateClusterTemplateDefault struct {
	_statusCode int
}

// Code gets the status code for the delete private cluster template default response
func (o *DeletePrivateClusterTemplateDefault) Code() int {
	return o._statusCode
}

func (o *DeletePrivateClusterTemplateDefault) Error() string {
	return fmt.Sprintf("[DELETE /clustertemplates/user/{name}][%d] deletePrivateClusterTemplate default ", o._statusCode)
}

func (o *DeletePrivateClusterTemplateDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}