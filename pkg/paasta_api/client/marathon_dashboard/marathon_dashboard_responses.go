// Code generated by go-swagger; DO NOT EDIT.

package marathon_dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Yelp/paasta-tools-go/pkg/paasta_api/models"
)

// MarathonDashboardReader is a Reader for the MarathonDashboard structure.
type MarathonDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MarathonDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMarathonDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMarathonDashboardOK creates a MarathonDashboardOK with default headers values
func NewMarathonDashboardOK() *MarathonDashboardOK {
	return &MarathonDashboardOK{}
}

/*MarathonDashboardOK handles this case with default header values.

List of service instances and information on their Marathon shard
*/
type MarathonDashboardOK struct {
	Payload models.MarathonDashboard
}

func (o *MarathonDashboardOK) Error() string {
	return fmt.Sprintf("[GET /marathon_dashboard][%d] marathonDashboardOK  %+v", 200, o.Payload)
}

func (o *MarathonDashboardOK) GetPayload() models.MarathonDashboard {
	return o.Payload
}

func (o *MarathonDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
