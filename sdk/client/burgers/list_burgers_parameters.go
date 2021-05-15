// Code generated by go-swagger; DO NOT EDIT.

package burgers

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

// NewListBurgersParams creates a new ListBurgersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListBurgersParams() *ListBurgersParams {
	return &ListBurgersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListBurgersParamsWithTimeout creates a new ListBurgersParams object
// with the ability to set a timeout on a request.
func NewListBurgersParamsWithTimeout(timeout time.Duration) *ListBurgersParams {
	return &ListBurgersParams{
		timeout: timeout,
	}
}

// NewListBurgersParamsWithContext creates a new ListBurgersParams object
// with the ability to set a context for a request.
func NewListBurgersParamsWithContext(ctx context.Context) *ListBurgersParams {
	return &ListBurgersParams{
		Context: ctx,
	}
}

// NewListBurgersParamsWithHTTPClient creates a new ListBurgersParams object
// with the ability to set a custom HTTPClient for a request.
func NewListBurgersParamsWithHTTPClient(client *http.Client) *ListBurgersParams {
	return &ListBurgersParams{
		HTTPClient: client,
	}
}

/* ListBurgersParams contains all the parameters to send to the API endpoint
   for the list burgers operation.

   Typically these are written to a http.Request.
*/
type ListBurgersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list burgers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBurgersParams) WithDefaults() *ListBurgersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list burgers params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListBurgersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list burgers params
func (o *ListBurgersParams) WithTimeout(timeout time.Duration) *ListBurgersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list burgers params
func (o *ListBurgersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list burgers params
func (o *ListBurgersParams) WithContext(ctx context.Context) *ListBurgersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list burgers params
func (o *ListBurgersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list burgers params
func (o *ListBurgersParams) WithHTTPClient(client *http.Client) *ListBurgersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list burgers params
func (o *ListBurgersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListBurgersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}