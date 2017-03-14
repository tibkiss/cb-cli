package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// DeleteLdapAccountNameReader is a Reader for the DeleteLdapAccountName structure.
type DeleteLdapAccountNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteLdapAccountNameReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewDeleteLdapAccountNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteLdapAccountNameDefault creates a DeleteLdapAccountNameDefault with default headers values
func NewDeleteLdapAccountNameDefault(code int) *DeleteLdapAccountNameDefault {
	return &DeleteLdapAccountNameDefault{
		_statusCode: code,
	}
}

/*DeleteLdapAccountNameDefault handles this case with default header values.

successful operation
*/
type DeleteLdapAccountNameDefault struct {
	_statusCode int
}

// Code gets the status code for the delete ldap account name default response
func (o *DeleteLdapAccountNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteLdapAccountNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /ldap/account/{name}][%d] DeleteLdapAccountName default ", o._statusCode)
}

func (o *DeleteLdapAccountNameDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}