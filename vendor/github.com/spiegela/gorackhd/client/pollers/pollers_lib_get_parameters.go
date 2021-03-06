package pollers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPollersLibGetParams creates a new PollersLibGetParams object
// with the default values initialized.
func NewPollersLibGetParams() *PollersLibGetParams {

	return &PollersLibGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPollersLibGetParamsWithTimeout creates a new PollersLibGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPollersLibGetParamsWithTimeout(timeout time.Duration) *PollersLibGetParams {

	return &PollersLibGetParams{

		timeout: timeout,
	}
}

// NewPollersLibGetParamsWithContext creates a new PollersLibGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPollersLibGetParamsWithContext(ctx context.Context) *PollersLibGetParams {

	return &PollersLibGetParams{

		Context: ctx,
	}
}

// NewPollersLibGetParamsWithHTTPClient creates a new PollersLibGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPollersLibGetParamsWithHTTPClient(client *http.Client) *PollersLibGetParams {

	return &PollersLibGetParams{
		HTTPClient: client,
	}
}

/*PollersLibGetParams contains all the parameters to send to the API endpoint
for the pollers lib get operation typically these are written to a http.Request
*/
type PollersLibGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pollers lib get params
func (o *PollersLibGetParams) WithTimeout(timeout time.Duration) *PollersLibGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pollers lib get params
func (o *PollersLibGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pollers lib get params
func (o *PollersLibGetParams) WithContext(ctx context.Context) *PollersLibGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pollers lib get params
func (o *PollersLibGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pollers lib get params
func (o *PollersLibGetParams) WithHTTPClient(client *http.Client) *PollersLibGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pollers lib get params
func (o *PollersLibGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PollersLibGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
