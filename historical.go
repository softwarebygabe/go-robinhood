package robinhood

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

// HistQueryOptions are the options taken by the "historicals" API
type HistQueryOptions struct {
	Interval string
	Bounds   string
	Span     string
}

// DefaultHistoricalOptions returns default options for the "historicals" API
// Defaults are interval=5minute, bounds=trading, span=day
func DefaultHistoricalOptions() HistQueryOptions {
	return HistQueryOptions{
		Interval: "5minute",
		Bounds:   "trading",
		Span:     "day",
	}
}

// GetHistoricals returns the historical candle data associated with a given instrument.
// It utilizes the "/marketdata/historicals" API.
func (c *Client) GetHistoricals(instrument Instrument) (Historical, error) {
	defaultOpts := DefaultHistoricalOptions()
	queryString := "?interval=" + defaultOpts.Interval + "&bounds=" + defaultOpts.Bounds + "&span=" + defaultOpts.Span
	var h Historical
	err := c.GetAndDecode(EPHistoricals+instrument.ID+"/"+queryString, &h)
	return h, err
}
