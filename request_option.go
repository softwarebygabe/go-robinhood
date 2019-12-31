package robinhood

import (
	"net/http"

	"github.com/blend/go-sdk/webutil"
)

// ApplyRequestOptions ...
func ApplyRequestOptions(req *http.Request, options ...webutil.RequestOption) error {
	for _, opt := range options {
		if err := opt(req); err != nil {
			return err
		}
	}
	return nil
}
