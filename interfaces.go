package robinhood

import "github.com/blend/go-sdk/webutil"

// AccountGetter can retrieve account data
type AccountGetter interface {
	GetAccounts() ([]Account, error)
}

// InstrumentGetter can retrieve instrument data
type InstrumentGetter interface {
	GetInstrument(instURL string) (*Instrument, error)
	GetInstrumentForSymbol(sym string) (*Instrument, error)
}

// OrderManager can make order operations
type OrderManager interface {
	Order(i *Instrument, o OrderOpts) (*OrderOutput, error)
	UpdateOrder(o *OrderOutput) error
	CancelOrder(o *OrderOutput) error
	RecentOrders() ([]OrderOutput, error)
}

// FundamentalsGetter can retrieve fundamentals data
type FundamentalsGetter interface {
	GetFundamentals(stocks ...string) ([]Fundamental, error)
}

// QuoteGetter can retreive current quote data
type QuoteGetter interface {
	GetQuote(stocks ...string) ([]Quote, error)
}

// HistoricalGetter can retrieve historical data
type HistoricalGetter interface {
	GetHistoricals(instrument Instrument, options ...webutil.RequestOption) (Historical, error)
}

// PositionsGetter can retrieve position data
type PositionsGetter interface {
	GetPositions(a Account) ([]Position, error)
	GetPositionsParams(a Account, p PositionParams) ([]Position, error)
}

// WatchlistGetter can retrieve watchlist data
type WatchlistGetter interface {
	GetWatchlists() ([]Watchlist, error)
}
