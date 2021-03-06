// Code generated by go-swagger; DO NOT EDIT.

package credentials

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

// NewListCredentialsParams creates a new ListCredentialsParams object
// with the default values initialized.
func NewListCredentialsParams() *ListCredentialsParams {
	var ()
	return &ListCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListCredentialsParamsWithTimeout creates a new ListCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListCredentialsParamsWithTimeout(timeout time.Duration) *ListCredentialsParams {
	var ()
	return &ListCredentialsParams{

		timeout: timeout,
	}
}

// NewListCredentialsParamsWithContext creates a new ListCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListCredentialsParamsWithContext(ctx context.Context) *ListCredentialsParams {
	var ()
	return &ListCredentialsParams{

		Context: ctx,
	}
}

// NewListCredentialsParamsWithHTTPClient creates a new ListCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListCredentialsParamsWithHTTPClient(client *http.Client) *ListCredentialsParams {
	var ()
	return &ListCredentialsParams{
		HTTPClient: client,
	}
}

/*ListCredentialsParams contains all the parameters to send to the API endpoint
for the list credentials operation typically these are written to a http.Request
*/
type ListCredentialsParams struct {

	/*Datacenter*/
	Datacenter *string
	/*ProviderName*/
	ProviderName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list credentials params
func (o *ListCredentialsParams) WithTimeout(timeout time.Duration) *ListCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list credentials params
func (o *ListCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list credentials params
func (o *ListCredentialsParams) WithContext(ctx context.Context) *ListCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list credentials params
func (o *ListCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list credentials params
func (o *ListCredentialsParams) WithHTTPClient(client *http.Client) *ListCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list credentials params
func (o *ListCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatacenter adds the datacenter to the list credentials params
func (o *ListCredentialsParams) WithDatacenter(datacenter *string) *ListCredentialsParams {
	o.SetDatacenter(datacenter)
	return o
}

// SetDatacenter adds the datacenter to the list credentials params
func (o *ListCredentialsParams) SetDatacenter(datacenter *string) {
	o.Datacenter = datacenter
}

// WithProviderName adds the providerName to the list credentials params
func (o *ListCredentialsParams) WithProviderName(providerName string) *ListCredentialsParams {
	o.SetProviderName(providerName)
	return o
}

// SetProviderName adds the providerName to the list credentials params
func (o *ListCredentialsParams) SetProviderName(providerName string) {
	o.ProviderName = providerName
}

// WriteToRequest writes these params to a swagger request
func (o *ListCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Datacenter != nil {

		// query param datacenter
		var qrDatacenter string
		if o.Datacenter != nil {
			qrDatacenter = *o.Datacenter
		}
		qDatacenter := qrDatacenter
		if qDatacenter != "" {
			if err := r.SetQueryParam("datacenter", qDatacenter); err != nil {
				return err
			}
		}

	}

	// path param provider_name
	if err := r.SetPathParam("provider_name", o.ProviderName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
