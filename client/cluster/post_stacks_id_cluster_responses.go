package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// PostStacksIDClusterReader is a Reader for the PostStacksIDCluster structure.
type PostStacksIDClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *PostStacksIDClusterReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewPostStacksIDClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPostStacksIDClusterDefault creates a PostStacksIDClusterDefault with default headers values
func NewPostStacksIDClusterDefault(code int) *PostStacksIDClusterDefault {
	return &PostStacksIDClusterDefault{
		_statusCode: code,
	}
}

/*PostStacksIDClusterDefault handles this case with default header values.

successful operation
*/
type PostStacksIDClusterDefault struct {
	_statusCode int
}

// Code gets the status code for the post stacks ID cluster default response
func (o *PostStacksIDClusterDefault) Code() int {
	return o._statusCode
}

func (o *PostStacksIDClusterDefault) Error() string {
	return fmt.Sprintf("[POST /stacks/{id}/cluster][%d] PostStacksIDCluster default ", o._statusCode)
}

func (o *PostStacksIDClusterDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}