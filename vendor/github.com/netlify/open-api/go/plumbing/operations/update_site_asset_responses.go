// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// UpdateSiteAssetReader is a Reader for the UpdateSiteAsset structure.
type UpdateSiteAssetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSiteAssetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateSiteAssetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateSiteAssetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSiteAssetOK creates a UpdateSiteAssetOK with default headers values
func NewUpdateSiteAssetOK() *UpdateSiteAssetOK {
	return &UpdateSiteAssetOK{}
}

/*UpdateSiteAssetOK handles this case with default header values.

Updated
*/
type UpdateSiteAssetOK struct {
	Payload *models.Asset
}

func (o *UpdateSiteAssetOK) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/assets/{asset_id}][%d] updateSiteAssetOK  %+v", 200, o.Payload)
}

func (o *UpdateSiteAssetOK) GetPayload() *models.Asset {
	return o.Payload
}

func (o *UpdateSiteAssetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Asset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateSiteAssetDefault creates a UpdateSiteAssetDefault with default headers values
func NewUpdateSiteAssetDefault(code int) *UpdateSiteAssetDefault {
	return &UpdateSiteAssetDefault{
		_statusCode: code,
	}
}

/*UpdateSiteAssetDefault handles this case with default header values.

error
*/
type UpdateSiteAssetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the update site asset default response
func (o *UpdateSiteAssetDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSiteAssetDefault) Error() string {
	return fmt.Sprintf("[PUT /sites/{site_id}/assets/{asset_id}][%d] updateSiteAsset default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSiteAssetDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateSiteAssetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
