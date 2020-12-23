// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewShowServiceManifestParams creates a new ShowServiceManifestParams object
// with the default values initialized.
func NewShowServiceManifestParams() *ShowServiceManifestParams {
	var ()
	return &ShowServiceManifestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewShowServiceManifestParamsWithTimeout creates a new ShowServiceManifestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewShowServiceManifestParamsWithTimeout(timeout time.Duration) *ShowServiceManifestParams {
	var ()
	return &ShowServiceManifestParams{

		timeout: timeout,
	}
}

// NewShowServiceManifestParamsWithContext creates a new ShowServiceManifestParams object
// with the default values initialized, and the ability to set a context for a request
func NewShowServiceManifestParamsWithContext(ctx context.Context) *ShowServiceManifestParams {
	var ()
	return &ShowServiceManifestParams{

		Context: ctx,
	}
}

// NewShowServiceManifestParamsWithHTTPClient creates a new ShowServiceManifestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewShowServiceManifestParamsWithHTTPClient(client *http.Client) *ShowServiceManifestParams {
	var ()
	return &ShowServiceManifestParams{
		HTTPClient: client,
	}
}

/*ShowServiceManifestParams contains all the parameters to send to the API endpoint
for the show service manifest operation typically these are written to a http.Request
*/
type ShowServiceManifestParams struct {

	/*AddonName*/
	AddonName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the show service manifest params
func (o *ShowServiceManifestParams) WithTimeout(timeout time.Duration) *ShowServiceManifestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the show service manifest params
func (o *ShowServiceManifestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the show service manifest params
func (o *ShowServiceManifestParams) WithContext(ctx context.Context) *ShowServiceManifestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the show service manifest params
func (o *ShowServiceManifestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the show service manifest params
func (o *ShowServiceManifestParams) WithHTTPClient(client *http.Client) *ShowServiceManifestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the show service manifest params
func (o *ShowServiceManifestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddonName adds the addonName to the show service manifest params
func (o *ShowServiceManifestParams) WithAddonName(addonName string) *ShowServiceManifestParams {
	o.SetAddonName(addonName)
	return o
}

// SetAddonName adds the addonName to the show service manifest params
func (o *ShowServiceManifestParams) SetAddonName(addonName string) {
	o.AddonName = addonName
}

// WriteToRequest writes these params to a swagger request
func (o *ShowServiceManifestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param addonName
	if err := r.SetPathParam("addonName", o.AddonName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
