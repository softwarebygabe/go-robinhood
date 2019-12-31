package robinhood

import (
	"github.com/blend/go-sdk/webutil"
)

/*
{
	"begins_at":"2019-08-12T13:00:00Z",
	"open_price":"132.910000",
	"close_price":"132.910000",
	"high_price":"132.910000",
	"low_price":"132.910000",
	"volume":0,
	"session":"pre",
	"interpolated":true
}
*/

// Candle represents the metrics associated with an instrument and
// its price over a given time period
type Candle struct {
	BeginsAt     string `json:"begins_at"`
	OpenPrice    string `json:"open_price"`
	ClosePrice   string `json:"close_price"`
	HighPrice    string `json:"high_price"`
	LowPrice     string `json:"low_price"`
	Volume       int    `json:"volume"`
	Session      string `json:"session"`
	Interpolated bool   `json:"interpolated"`
}

/*
{
	"quote":"https://api.robinhood.com/quotes/d57904fb-55fe-4e2b-97f7-34ef2e0729ec/",
	"symbol":"OKTA",
	"interval":"5minute",
	"span":"day",
	"bounds":"trading",
	"previous_close_price":"134.200000",
	"previous_close_time":"2019-08-09T20:00:00Z",
	"open_price":"132.910000",
	"open_time":"2019-08-12T13:00:00Z",
	"instrument":"https://api.robinhood.com/instruments/d57904fb-55fe-4e2b-97f7-34ef2e0729ec/",
	"historicals":[]
}
*/

// Historical represents the response of the "historicals" API
type Historical struct {
	Quote              string   `json:"quote"`
	Symbol             string   `json:"symbol"`
	Interval           string   `json:"interval"`
	Span               string   `json:"span"`
	Bounds             string   `json:"bounds"`
	PreviousClosePrice string   `json:"previous_close_price"`
	PreviousCloseTime  string   `json:"previous_close_time"`
	OpenPrice          string   `json:"open_price"`
	OpenTime           string   `json:"open_time"`
	Instrument         string   `json:"instrument"`
	Historicals        []Candle `json:"historicals"`
}

// GetHistoricals returns the historical candle data associated with a given instrument.
// It utilizes the "/marketdata/historicals" API.
func (c *Client) GetHistoricals(instrument Instrument, options ...webutil.RequestOption) (Historical, error) {
	var h Historical
	defaultOptions := []webutil.RequestOption{
		OptHistoricalsInterval5Minute(),
		OptHistoricalsSpanDay(),
		OptHistoricalsBoundsTrading(),
	}
	allOptions := append(defaultOptions, options...)
	req, err := c.NewRequest(
		"GET",
		EPHistoricals+instrument.ID,
		allOptions...,
	)
	if err != nil {
		return h, err
	}
	err = c.DoAndDecode(req, &h)
	return h, err
}

// OptHistoricalsInterval5Minute sets the `interval` to `5minute` in the request for historicals
func OptHistoricalsInterval5Minute() webutil.RequestOption {
	return webutil.OptQueryValue("interval", "5minute")
}

// OptHistoricalsBoundsTrading sets the `bounds` to `trading` in the request for historicals
func OptHistoricalsBoundsTrading() webutil.RequestOption {
	return webutil.OptQueryValue("bounds", "trading")
}

// OptHistoricalsSpanDay sets the `span` to `day` in the request for historicals
func OptHistoricalsSpanDay() webutil.RequestOption {
	return webutil.OptQueryValue("span", "day")
}
