// Code generated by go-swagger; DO NOT EDIT.

package v1imagecatalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePublicImageCatalogByNameReader is a Reader for the DeletePublicImageCatalogByName structure.
type DeletePublicImageCatalogByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePublicImageCatalogByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {

	result := NewDeletePublicImageCatalogByNameDefault(response.Code())
	if err := result.readResponse(response, consumer, o.formats); err != nil {
		return nil, err
	}
	if response.Code()/100 == 2 {
		return result, nil
	}
	return nil, result

}

// NewDeletePublicImageCatalogByNameDefault creates a DeletePublicImageCatalogByNameDefault with default headers values
func NewDeletePublicImageCatalogByNameDefault(code int) *DeletePublicImageCatalogByNameDefault {
	return &DeletePublicImageCatalogByNameDefault{
		_statusCode: code,
	}
}

/*DeletePublicImageCatalogByNameDefault handles this case with default header values.

successful operation
*/
type DeletePublicImageCatalogByNameDefault struct {
	_statusCode int
}

// Code gets the status code for the delete public image catalog by name default response
func (o *DeletePublicImageCatalogByNameDefault) Code() int {
	return o._statusCode
}

func (o *DeletePublicImageCatalogByNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/imagecatalogs/account/{name}][%d] deletePublicImageCatalogByName default ", o._statusCode)
}

func (o *DeletePublicImageCatalogByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
