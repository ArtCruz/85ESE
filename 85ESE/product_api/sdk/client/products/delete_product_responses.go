// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"product_api/sdk/models"
)

// DeleteProductReader is a Reader for the DeleteProduct structure.
type DeleteProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDeleteProductCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewDeleteProductNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewDeleteProductNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /products/{id}] deleteProduct", response, response.Code())
	}
}

// NewDeleteProductCreated creates a DeleteProductCreated with default headers values
func NewDeleteProductCreated() *DeleteProductCreated {
	return &DeleteProductCreated{}
}

/*
DeleteProductCreated describes a response with status code 201, with default header values.

No content is returned by this API endpoint
*/
type DeleteProductCreated struct {
}

// IsSuccess returns true when this delete product created response has a 2xx status code
func (o *DeleteProductCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete product created response has a 3xx status code
func (o *DeleteProductCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete product created response has a 4xx status code
func (o *DeleteProductCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete product created response has a 5xx status code
func (o *DeleteProductCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this delete product created response a status code equal to that given
func (o *DeleteProductCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the delete product created response
func (o *DeleteProductCreated) Code() int {
	return 201
}

func (o *DeleteProductCreated) Error() string {
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductCreated", 201)
}

func (o *DeleteProductCreated) String() string {
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductCreated", 201)
}

func (o *DeleteProductCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteProductNotFound creates a DeleteProductNotFound with default headers values
func NewDeleteProductNotFound() *DeleteProductNotFound {
	return &DeleteProductNotFound{}
}

/*
DeleteProductNotFound describes a response with status code 404, with default header values.

Generic error message returned as a string
*/
type DeleteProductNotFound struct {
	Payload *models.GenericError
}

// IsSuccess returns true when this delete product not found response has a 2xx status code
func (o *DeleteProductNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete product not found response has a 3xx status code
func (o *DeleteProductNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete product not found response has a 4xx status code
func (o *DeleteProductNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete product not found response has a 5xx status code
func (o *DeleteProductNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete product not found response a status code equal to that given
func (o *DeleteProductNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete product not found response
func (o *DeleteProductNotFound) Code() int {
	return 404
}

func (o *DeleteProductNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductNotFound %s", 404, payload)
}

func (o *DeleteProductNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductNotFound %s", 404, payload)
}

func (o *DeleteProductNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteProductNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteProductNotImplemented creates a DeleteProductNotImplemented with default headers values
func NewDeleteProductNotImplemented() *DeleteProductNotImplemented {
	return &DeleteProductNotImplemented{}
}

/*
DeleteProductNotImplemented describes a response with status code 501, with default header values.

Generic error message returned as a string
*/
type DeleteProductNotImplemented struct {
	Payload *models.GenericError
}

// IsSuccess returns true when this delete product not implemented response has a 2xx status code
func (o *DeleteProductNotImplemented) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete product not implemented response has a 3xx status code
func (o *DeleteProductNotImplemented) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete product not implemented response has a 4xx status code
func (o *DeleteProductNotImplemented) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete product not implemented response has a 5xx status code
func (o *DeleteProductNotImplemented) IsServerError() bool {
	return true
}

// IsCode returns true when this delete product not implemented response a status code equal to that given
func (o *DeleteProductNotImplemented) IsCode(code int) bool {
	return code == 501
}

// Code gets the status code for the delete product not implemented response
func (o *DeleteProductNotImplemented) Code() int {
	return 501
}

func (o *DeleteProductNotImplemented) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductNotImplemented %s", 501, payload)
}

func (o *DeleteProductNotImplemented) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /products/{id}][%d] deleteProductNotImplemented %s", 501, payload)
}

func (o *DeleteProductNotImplemented) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *DeleteProductNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
