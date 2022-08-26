// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-swagger/go-swagger/examples/todo-list/models"
)

// FindReader is a Reader for the Find structure.
type FindReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindOK creates a FindOK with default headers values
func NewFindOK() *FindOK {
	return &FindOK{}
}

/*
	FindOK describes a response with status code 200, with default header values.

OK
*/
type FindOK struct {
	Payload []*models.Item
}

// IsSuccess returns true when this find o k response returns a 2xx status code
func (o *FindOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find o k response returns a 3xx status code
func (o *FindOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find o k response returns a 4xx status code
func (o *FindOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find o k response returns a 5xx status code
func (o *FindOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find o k response returns a 4xx status code
func (o *FindOK) IsCode(code int) bool {
	return code == 200
}

func (o *FindOK) Error() string {
	return fmt.Sprintf("[GET /][%d] findOK  %+v", 200, o.Payload)
}

func (o *FindOK) String() string {
	return fmt.Sprintf("[GET /][%d] findOK  %+v", 200, o.Payload)
}

func (o *FindOK) GetPayload() []*models.Item {
	return o.Payload
}

func (o *FindOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindDefault creates a FindDefault with default headers values
func NewFindDefault(code int) *FindDefault {
	return &FindDefault{
		_statusCode: code,
	}
}

/*
	FindDefault describes a response with status code -1, with default header values.

error
*/
type FindDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the find default response
func (o *FindDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this find default response returns a 2xx status code
func (o *FindDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this find default response returns a 3xx status code
func (o *FindDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this find default response returns a 4xx status code
func (o *FindDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this find default response returns a 5xx status code
func (o *FindDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this find default response returns a 4xx status code
func (o *FindDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FindDefault) Error() string {
	return fmt.Sprintf("[GET /][%d] find default  %+v", o._statusCode, o.Payload)
}

func (o *FindDefault) String() string {
	return fmt.Sprintf("[GET /][%d] find default  %+v", o._statusCode, o.Payload)
}

func (o *FindDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *FindDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
