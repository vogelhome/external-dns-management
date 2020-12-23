// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUploadDeployFunctionParams creates a new UploadDeployFunctionParams object
// with the default values initialized.
func NewUploadDeployFunctionParams() *UploadDeployFunctionParams {
	var ()
	return &UploadDeployFunctionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadDeployFunctionParamsWithTimeout creates a new UploadDeployFunctionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadDeployFunctionParamsWithTimeout(timeout time.Duration) *UploadDeployFunctionParams {
	var ()
	return &UploadDeployFunctionParams{

		timeout: timeout,
	}
}

// NewUploadDeployFunctionParamsWithContext creates a new UploadDeployFunctionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadDeployFunctionParamsWithContext(ctx context.Context) *UploadDeployFunctionParams {
	var ()
	return &UploadDeployFunctionParams{

		Context: ctx,
	}
}

// NewUploadDeployFunctionParamsWithHTTPClient creates a new UploadDeployFunctionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadDeployFunctionParamsWithHTTPClient(client *http.Client) *UploadDeployFunctionParams {
	var ()
	return &UploadDeployFunctionParams{
		HTTPClient: client,
	}
}

/*UploadDeployFunctionParams contains all the parameters to send to the API endpoint
for the upload deploy function operation typically these are written to a http.Request
*/
type UploadDeployFunctionParams struct {

	/*DeployID*/
	DeployID string
	/*FileBody*/
	FileBody io.ReadCloser
	/*Name*/
	Name string
	/*Runtime*/
	Runtime *string
	/*Size*/
	Size *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upload deploy function params
func (o *UploadDeployFunctionParams) WithTimeout(timeout time.Duration) *UploadDeployFunctionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload deploy function params
func (o *UploadDeployFunctionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload deploy function params
func (o *UploadDeployFunctionParams) WithContext(ctx context.Context) *UploadDeployFunctionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload deploy function params
func (o *UploadDeployFunctionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload deploy function params
func (o *UploadDeployFunctionParams) WithHTTPClient(client *http.Client) *UploadDeployFunctionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload deploy function params
func (o *UploadDeployFunctionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeployID adds the deployID to the upload deploy function params
func (o *UploadDeployFunctionParams) WithDeployID(deployID string) *UploadDeployFunctionParams {
	o.SetDeployID(deployID)
	return o
}

// SetDeployID adds the deployId to the upload deploy function params
func (o *UploadDeployFunctionParams) SetDeployID(deployID string) {
	o.DeployID = deployID
}

// WithFileBody adds the fileBody to the upload deploy function params
func (o *UploadDeployFunctionParams) WithFileBody(fileBody io.ReadCloser) *UploadDeployFunctionParams {
	o.SetFileBody(fileBody)
	return o
}

// SetFileBody adds the fileBody to the upload deploy function params
func (o *UploadDeployFunctionParams) SetFileBody(fileBody io.ReadCloser) {
	o.FileBody = fileBody
}

// WithName adds the name to the upload deploy function params
func (o *UploadDeployFunctionParams) WithName(name string) *UploadDeployFunctionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the upload deploy function params
func (o *UploadDeployFunctionParams) SetName(name string) {
	o.Name = name
}

// WithRuntime adds the runtime to the upload deploy function params
func (o *UploadDeployFunctionParams) WithRuntime(runtime *string) *UploadDeployFunctionParams {
	o.SetRuntime(runtime)
	return o
}

// SetRuntime adds the runtime to the upload deploy function params
func (o *UploadDeployFunctionParams) SetRuntime(runtime *string) {
	o.Runtime = runtime
}

// WithSize adds the size to the upload deploy function params
func (o *UploadDeployFunctionParams) WithSize(size *int64) *UploadDeployFunctionParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the upload deploy function params
func (o *UploadDeployFunctionParams) SetSize(size *int64) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *UploadDeployFunctionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deploy_id
	if err := r.SetPathParam("deploy_id", o.DeployID); err != nil {
		return err
	}

	if o.FileBody != nil {
		if err := r.SetBodyParam(o.FileBody); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if o.Runtime != nil {

		// query param runtime
		var qrRuntime string
		if o.Runtime != nil {
			qrRuntime = *o.Runtime
		}
		qRuntime := qrRuntime
		if qRuntime != "" {
			if err := r.SetQueryParam("runtime", qRuntime); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize int64
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt64(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
