// Code generated by go-swagger; DO NOT EDIT.

package v1util

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// TestLdapConnectionByIDUtilReader is a Reader for the TestLdapConnectionByIDUtil structure.
type TestLdapConnectionByIDUtilReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestLdapConnectionByIDUtilReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewTestLdapConnectionByIDUtilOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTestLdapConnectionByIDUtilOK creates a TestLdapConnectionByIDUtilOK with default headers values
func NewTestLdapConnectionByIDUtilOK() *TestLdapConnectionByIDUtilOK {
	return &TestLdapConnectionByIDUtilOK{}
}

/*TestLdapConnectionByIDUtilOK handles this case with default header values.

successful operation
*/
type TestLdapConnectionByIDUtilOK struct {
	Payload *models_cloudbreak.RdsTestResult
}

func (o *TestLdapConnectionByIDUtilOK) Error() string {
	return fmt.Sprintf("[GET /v1/util/ldap/{id}][%d] testLdapConnectionByIdUtilOK  %+v", 200, o.Payload)
}

func (o *TestLdapConnectionByIDUtilOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.RdsTestResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
