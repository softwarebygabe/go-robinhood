# Robinhood API Golang Client

This library is forked from [Andrew Stuart's](https://github.com/andrewstuart/go-robinhood) and builds upon the work done there, namely
the oauth2 implementation and the characterization of a lot of the API endpoints and response types.

My implementation aims to add more API functionality as it becomes available from Robinhood, as well as to improve the usability
and flexibility of the library and its functions by implementing the [Functional Options](https://github.com/tmrts/go-patterns/blob/master/idiom/functional-options.md) pattern, and utilizing all-purpose libraries like [Blend's Go SDK](https://github.com/blend/go-sdk) for
things like logging and webutils.

## General usage

```go
client, _ := robinhood.NewClient(
  robinhood.NewOAuth("username", "password"),
)

instrument, _ := robinhood.GetInstrumentForSymbol("SPY")

historicalData, _ := robinhood.GetHistoricals(
  instrument,
  robinhood.OptHistoricalsInterval5Minute(),
  robinhood.OptHistoricalsSpanDay(),
  robinhood.OptHistoricalsBoundsTrading(),
)

// Do something with the data :)
```

## Client Options

```go
import (
  "github.com/blend/go-sdk/logger
)

logger := logger.MustNew()

client, _ := robinhood.NewClient(
  robinhood.NewOAuth("username", "password"),
  robinhood.OptClientLog(logger),
)

// now client will use the provided logger...
```
