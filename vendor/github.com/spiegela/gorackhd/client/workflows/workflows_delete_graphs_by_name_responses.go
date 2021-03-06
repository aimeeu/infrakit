package workflows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// WorkflowsDeleteGraphsByNameReader is a Reader for the WorkflowsDeleteGraphsByName structure.
type WorkflowsDeleteGraphsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkflowsDeleteGraphsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewWorkflowsDeleteGraphsByNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewWorkflowsDeleteGraphsByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewWorkflowsDeleteGraphsByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWorkflowsDeleteGraphsByNameNoContent creates a WorkflowsDeleteGraphsByNameNoContent with default headers values
func NewWorkflowsDeleteGraphsByNameNoContent() *WorkflowsDeleteGraphsByNameNoContent {
	return &WorkflowsDeleteGraphsByNameNoContent{}
}

/*WorkflowsDeleteGraphsByNameNoContent handles this case with default header values.

Successfully deleted the specified workflow graph
*/
type WorkflowsDeleteGraphsByNameNoContent struct {
	Payload WorkflowsDeleteGraphsByNameNoContentBody
}

func (o *WorkflowsDeleteGraphsByNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /workflows/graphs/{injectableName}][%d] workflowsDeleteGraphsByNameNoContent  %+v", 204, o.Payload)
}

func (o *WorkflowsDeleteGraphsByNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkflowsDeleteGraphsByNameNotFound creates a WorkflowsDeleteGraphsByNameNotFound with default headers values
func NewWorkflowsDeleteGraphsByNameNotFound() *WorkflowsDeleteGraphsByNameNotFound {
	return &WorkflowsDeleteGraphsByNameNotFound{}
}

/*WorkflowsDeleteGraphsByNameNotFound handles this case with default header values.

The graph with the specified injectable name was not found
*/
type WorkflowsDeleteGraphsByNameNotFound struct {
	Payload *models.Error
}

func (o *WorkflowsDeleteGraphsByNameNotFound) Error() string {
	return fmt.Sprintf("[DELETE /workflows/graphs/{injectableName}][%d] workflowsDeleteGraphsByNameNotFound  %+v", 404, o.Payload)
}

func (o *WorkflowsDeleteGraphsByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkflowsDeleteGraphsByNameDefault creates a WorkflowsDeleteGraphsByNameDefault with default headers values
func NewWorkflowsDeleteGraphsByNameDefault(code int) *WorkflowsDeleteGraphsByNameDefault {
	return &WorkflowsDeleteGraphsByNameDefault{
		_statusCode: code,
	}
}

/*WorkflowsDeleteGraphsByNameDefault handles this case with default header values.

Unexpected error
*/
type WorkflowsDeleteGraphsByNameDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the workflows delete graphs by name default response
func (o *WorkflowsDeleteGraphsByNameDefault) Code() int {
	return o._statusCode
}

func (o *WorkflowsDeleteGraphsByNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /workflows/graphs/{injectableName}][%d] workflowsDeleteGraphsByName default  %+v", o._statusCode, o.Payload)
}

func (o *WorkflowsDeleteGraphsByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*WorkflowsDeleteGraphsByNameNoContentBody workflows delete graphs by name no content body
swagger:model WorkflowsDeleteGraphsByNameNoContentBody
*/
type WorkflowsDeleteGraphsByNameNoContentBody interface{}
