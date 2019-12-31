package robinhood

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/blend/go-sdk/logger"
	"github.com/blend/go-sdk/webutil"

	"golang.org/x/oauth2"
)

// Endpoints for the Robinhood API
const (
	EPBase         = "https://api.robinhood.com/"
	EPLogin        = EPBase + "oauth2/token/"
	EPAccounts     = EPBase + "accounts/"
	EPQuotes       = EPBase + "quotes/"
	EPPortfolios   = EPBase + "portfolios/"
	EPWatchlists   = EPBase + "watchlists/"
	EPInstruments  = EPBase + "instruments/"
	EPFundamentals = EPBase + "fundamentals/"
	EPOrders       = EPBase + "orders/"
	EPOptions      = EPBase + "options/"
	EPMarket       = EPBase + "marketdata/"
	EPOptionQuote  = EPMarket + "options/"
	EPHistoricals  = EPMarket + "historicals/"
)

// A Client is a helpful abstraction around some common metadata required for
// API operations.
type Client struct {
	Token   string
	Account *Account
	log     *logger.Logger
	*http.Client
}

// NewClient returns a client with the provided options set
func NewClient(httpClient *http.Client, opts ...ClientOption) (*Client, error) {
	c := &Client{
		Client: httpClient,
	}
	// set the client options
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// Dial returns a client given a TokenGetter. TokenGetter implementations are
// available in this package, including a Cookie-based cache.
func Dial(s oauth2.TokenSource) (*Client, error) {
	c := &Client{
		Client: oauth2.NewClient(context.Background(), s),
	}

	a, err := c.GetAccounts()
	if len(a) > 0 {
		c.Account = &a[0]
	}

	return c, err
}

// NewRequest creates a new http.Request with the given options.
func (c *Client) NewRequest(method, url string, options ...webutil.RequestOption) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	if err = ApplyRequestOptions(req, options...); err != nil {
		return nil, err
	}
	return req, nil
}

// GetAndDecode retrieves from the endpoint and unmarshals resulting json into
// the provided destination interface, which must be a pointer.
func (c *Client) GetAndDecode(url string, dest interface{}) error {
	req, err := c.NewRequest("GET", url)
	if err != nil {
		return err
	}

	return c.DoAndDecode(req, dest)
}

// ErrorMap encapsulates the helpful error messages returned by the API server
type ErrorMap map[string]interface{}

func (e ErrorMap) Error() string {
	es := make([]string, 0, len(e))
	for k, v := range e {
		es = append(es, fmt.Sprintf("%s: %q", k, v))
	}
	return "Error returned from API: " + strings.Join(es, ", ")
}

// DoAndDecode provides useful abstractions around common errors and decoding
// issues.
func (c *Client) DoAndDecode(req *http.Request, dest interface{}) error {
	c.MaybeInfof("making %s request to %s ...\n", req.Method, req.URL.RequestURI())
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode/100 != 2 {
		b := &bytes.Buffer{}
		var e ErrorMap
		err = json.NewDecoder(io.TeeReader(res.Body, b)).Decode(&e)
		if err != nil {
			return fmt.Errorf("got response %q and could not decode error body %q", res.Status, b.String())
		}
		return e
	}
	c.MaybeInfof("%d response received.\n", res.StatusCode)
	return json.NewDecoder(res.Body).Decode(dest)
}

// MaybeInfof logs if the client logger is set
func (c *Client) MaybeInfof(format string, args ...interface{}) {
	if c.log != nil {
		c.log.Infof(format, args...)
	}
}

// Meta holds metadata common to many RobinHood types.
type Meta struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
}
