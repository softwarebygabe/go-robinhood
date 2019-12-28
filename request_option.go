package robinhood

import "net/http"

// RequestOption is a modifier for a Request
type RequestOption func(*http.Request) error
