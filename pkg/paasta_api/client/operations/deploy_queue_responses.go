// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yelp/paasta-tools-go/pkg/paasta_api/models"
)

// DeployQueueReader is a Reader for the DeployQueue structure.
type DeployQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeployQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeployQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeployQueueOK creates a DeployQueueOK with default headers values
func NewDeployQueueOK() *DeployQueueOK {
	return &DeployQueueOK{}
}

/*DeployQueueOK handles this case with default header values.

Contents of deploy queue
*/
type DeployQueueOK struct {
	Payload *models.DeployQueue
}

func (o *DeployQueueOK) Error() string {
	return fmt.Sprintf("[GET /deploy_queue][%d] deployQueueOK  %+v", 200, o.Payload)
}

func (o *DeployQueueOK) GetPayload() *models.DeployQueue {
	return o.Payload
}

func (o *DeployQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeployQueue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
