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

// NewEstimateCoinBuyParams creates a new EstimateCoinBuyParams object
// with the default values initialized.
func NewEstimateCoinBuyParams() *EstimateCoinBuyParams {
	var ()
	return &EstimateCoinBuyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEstimateCoinBuyParamsWithTimeout creates a new EstimateCoinBuyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEstimateCoinBuyParamsWithTimeout(timeout time.Duration) *EstimateCoinBuyParams {
	var ()
	return &EstimateCoinBuyParams{

		timeout: timeout,
	}
}

// NewEstimateCoinBuyParamsWithContext creates a new EstimateCoinBuyParams object
// with the default values initialized, and the ability to set a context for a request
func NewEstimateCoinBuyParamsWithContext(ctx context.Context) *EstimateCoinBuyParams {
	var ()
	return &EstimateCoinBuyParams{

		Context: ctx,
	}
}

// NewEstimateCoinBuyParamsWithHTTPClient creates a new EstimateCoinBuyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEstimateCoinBuyParamsWithHTTPClient(client *http.Client) *EstimateCoinBuyParams {
	var ()
	return &EstimateCoinBuyParams{
		HTTPClient: client,
	}
}

/*EstimateCoinBuyParams contains all the parameters to send to the API endpoint
for the estimate coin buy operation typically these are written to a http.Request
*/
type EstimateCoinBuyParams struct {

	/*CoinIDToBuy*/
	CoinIDToBuy *uint64
	/*CoinIDToSell*/
	CoinIDToSell *uint64
	/*CoinToBuy*/
	CoinToBuy *string
	/*CoinToSell*/
	CoinToSell *string
	/*Height*/
	Height uint64
	/*ValueToBuy*/
	ValueToBuy string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the estimate coin buy params
func (o *EstimateCoinBuyParams) WithTimeout(timeout time.Duration) *EstimateCoinBuyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the estimate coin buy params
func (o *EstimateCoinBuyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the estimate coin buy params
func (o *EstimateCoinBuyParams) WithContext(ctx context.Context) *EstimateCoinBuyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the estimate coin buy params
func (o *EstimateCoinBuyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the estimate coin buy params
func (o *EstimateCoinBuyParams) WithHTTPClient(client *http.Client) *EstimateCoinBuyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the estimate coin buy params
func (o *EstimateCoinBuyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCoinIDToBuy adds the coinIDToBuy to the estimate coin buy params
func (o *EstimateCoinBuyParams) WithCoinIDToBuy(coinIDToBuy *uint64) *EstimateCoinBuyParams {
	o.SetCoinIDToBuy(coinIDToBuy)
	return o
}

// SetCoinIDToBuy adds the coinIdToBuy to the estimate coin buy params
func (o *EstimateCoinBuyParams) SetCoinIDToBuy(coinIDToBuy *uint64) {
	o.CoinIDToBuy = coinIDToBuy
}

// WithCoinIDToSell adds the coinIDToSell to the estimate coin buy params
func (o *EstimateCoinBuyParams) WithCoinIDToSell(coinIDToSell *uint64) *EstimateCoinBuyParams {
	o.SetCoinIDToSell(coinIDToSell)
	return o
}

// SetCoinIDToSell adds the coinIdToSell to the estimate coin buy params
func (o *EstimateCoinBuyParams) SetCoinIDToSell(coinIDToSell *uint64) {
	o.CoinIDToSell = coinIDToSell
}

// WithCoinToBuy adds the coinToBuy to the estimate coin buy params
func (o *EstimateCoinBuyParams) WithCoinToBuy(coinToBuy *string) *EstimateCoinBuyParams {
	o.SetCoinToBuy(coinToBuy)
	return o
}

// SetCoinToBuy adds the coinToBuy to the estimate coin buy params
func (o *EstimateCoinBuyParams) SetCoinToBuy(coinToBuy *string) {
	o.CoinToBuy = coinToBuy
}

// WithCoinToSell adds the coinToSell to the estimate coin buy params
func (o *EstimateCoinBuyParams) WithCoinToSell(coinToSell *string) *EstimateCoinBuyParams {
	o.SetCoinToSell(coinToSell)
	return o
}

// SetCoinToSell adds the coinToSell to the estimate coin buy params
func (o *EstimateCoinBuyParams) SetCoinToSell(coinToSell *string) {
	o.CoinToSell = coinToSell
}

// WithHeight adds the height to the estimate coin buy params
func (o *EstimateCoinBuyParams) WithHeight(height uint64) *EstimateCoinBuyParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the estimate coin buy params
func (o *EstimateCoinBuyParams) SetHeight(height uint64) {
	o.Height = height
}

// WithValueToBuy adds the valueToBuy to the estimate coin buy params
func (o *EstimateCoinBuyParams) WithValueToBuy(valueToBuy string) *EstimateCoinBuyParams {
	o.SetValueToBuy(valueToBuy)
	return o
}

// SetValueToBuy adds the valueToBuy to the estimate coin buy params
func (o *EstimateCoinBuyParams) SetValueToBuy(valueToBuy string) {
	o.ValueToBuy = valueToBuy
}

// WriteToRequest writes these params to a swagger request
func (o *EstimateCoinBuyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CoinIDToBuy != nil {

		// query param coin_id_to_buy
		var qrCoinIDToBuy uint64
		if o.CoinIDToBuy != nil {
			qrCoinIDToBuy = *o.CoinIDToBuy
		}
		qCoinIDToBuy := swag.FormatUint64(qrCoinIDToBuy)
		if qCoinIDToBuy != "" {
			if err := r.SetQueryParam("coin_id_to_buy", qCoinIDToBuy); err != nil {
				return err
			}
		}

	}

	if o.CoinIDToSell != nil {

		// query param coin_id_to_sell
		var qrCoinIDToSell uint64
		if o.CoinIDToSell != nil {
			qrCoinIDToSell = *o.CoinIDToSell
		}
		qCoinIDToSell := swag.FormatUint64(qrCoinIDToSell)
		if qCoinIDToSell != "" {
			if err := r.SetQueryParam("coin_id_to_sell", qCoinIDToSell); err != nil {
				return err
			}
		}

	}

	if o.CoinToBuy != nil {

		// query param coin_to_buy
		var qrCoinToBuy string
		if o.CoinToBuy != nil {
			qrCoinToBuy = *o.CoinToBuy
		}
		qCoinToBuy := qrCoinToBuy
		if qCoinToBuy != "" {
			if err := r.SetQueryParam("coin_to_buy", qCoinToBuy); err != nil {
				return err
			}
		}

	}

	if o.CoinToSell != nil {

		// query param coin_to_sell
		var qrCoinToSell string
		if o.CoinToSell != nil {
			qrCoinToSell = *o.CoinToSell
		}
		qCoinToSell := qrCoinToSell
		if qCoinToSell != "" {
			if err := r.SetQueryParam("coin_to_sell", qCoinToSell); err != nil {
				return err
			}
		}

	}

	// query param height
	qrHeight := o.Height
	qHeight := swag.FormatUint64(qrHeight)
	if qHeight != "" {
		if err := r.SetQueryParam("height", qHeight); err != nil {
			return err
		}
	}

	// query param value_to_buy
	qrValueToBuy := o.ValueToBuy
	qValueToBuy := qrValueToBuy
	if qValueToBuy != "" {
		if err := r.SetQueryParam("value_to_buy", qValueToBuy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
