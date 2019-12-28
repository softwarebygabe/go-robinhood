package robinhood

import (
	"github.com/blend/go-sdk/logger"
)

// ClientOption is a modifier for a Client
type ClientOption func(*Client) error

// OptClientLog sets the Client's logger to the one given
func OptClientLog(log *logger.Logger) ClientOption {
	return func(c *Client) error {
		if log != nil {
			c.log = log
		}
		return nil
	}
}
