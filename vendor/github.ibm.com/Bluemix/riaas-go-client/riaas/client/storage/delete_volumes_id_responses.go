// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// DeleteVolumesIDReader is a Reader for the DeleteVolumesID structure.
type DeleteVolumesIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteVolumesIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteVolumesIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewDeleteVolumesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteVolumesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteVolumesIDNoContent creates a DeleteVolumesIDNoContent with default headers values
func NewDeleteVolumesIDNoContent() *DeleteVolumesIDNoContent {
	return &DeleteVolumesIDNoContent{}
}

/*DeleteVolumesIDNoContent handles this case with default header values.

error
*/
type DeleteVolumesIDNoContent struct {
	Payload *models.Riaaserror
}

func (o *DeleteVolumesIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{id}][%d] deleteVolumesIdNoContent  %+v", 204, o.Payload)
}

func (o *DeleteVolumesIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVolumesIDNotFound creates a DeleteVolumesIDNotFound with default headers values
func NewDeleteVolumesIDNotFound() *DeleteVolumesIDNotFound {
	return &DeleteVolumesIDNotFound{}
}

/*DeleteVolumesIDNotFound handles this case with default header values.

error
*/
type DeleteVolumesIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *DeleteVolumesIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{id}][%d] deleteVolumesIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteVolumesIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteVolumesIDInternalServerError creates a DeleteVolumesIDInternalServerError with default headers values
func NewDeleteVolumesIDInternalServerError() *DeleteVolumesIDInternalServerError {
	return &DeleteVolumesIDInternalServerError{}
}

/*DeleteVolumesIDInternalServerError handles this case with default header values.

error
*/
type DeleteVolumesIDInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *DeleteVolumesIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /volumes/{id}][%d] deleteVolumesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteVolumesIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
