// Code generated by go-swagger; DO NOT EDIT.

package api_service

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
	"github.com/go-openapi/swag"
)

// NewEventsParams creates a new EventsParams object
// with the default values initialized.
func NewEventsParams() *EventsParams {
	var ()
	return &EventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEventsParamsWithTimeout creates a new EventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEventsParamsWithTimeout(timeout time.Duration) *EventsParams {
	var ()
	return &EventsParams{

		timeout: timeout,
	}
}

// NewEventsParamsWithContext creates a new EventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewEventsParamsWithContext(ctx context.Context) *EventsParams {
	var ()
	return &EventsParams{

		Context: ctx,
	}
}

// NewEventsParamsWithHTTPClient creates a new EventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEventsParamsWithHTTPClient(client *http.Client) *EventsParams {
	var ()
	return &EventsParams{
		HTTPClient: client,
	}
}

/*EventsParams contains all the parameters to send to the API endpoint
for the events operation typically these are written to a http.Request
*/
type EventsParams struct {

	/*Height*/
	Height string
	/*Search
	  Array of public keys of validators and wallet addresses of validators for filtering.

	*/
	Search []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the events params
func (o *EventsParams) WithTimeout(timeout time.Duration) *EventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the events params
func (o *EventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the events params
func (o *EventsParams) WithContext(ctx context.Context) *EventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the events params
func (o *EventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the events params
func (o *EventsParams) WithHTTPClient(client *http.Client) *EventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the events params
func (o *EventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the events params
func (o *EventsParams) WithHeight(height string) *EventsParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the events params
func (o *EventsParams) SetHeight(height string) {
	o.Height = height
}

// WithSearch adds the search to the events params
func (o *EventsParams) WithSearch(search []string) *EventsParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the events params
func (o *EventsParams) SetSearch(search []string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *EventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param height
	if err := r.SetPathParam("height", o.Height); err != nil {
		return err
	}

	valuesSearch := o.Search

	joinedSearch := swag.JoinByFormat(valuesSearch, "multi")
	// query array param search
	if err := r.SetQueryParam("search", joinedSearch...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
