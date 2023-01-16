package clients

import (
	"context"
)

type Client interface {
	ExchangeRate(ctx context.Context, assetBase string, assetQuote string) (float64, error)
}
