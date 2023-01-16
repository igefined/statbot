package coin

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
)

type (
	Rates struct {
		Time         string  `json:"time"`
		AssetIdBase  string  `json:"asset_id_base"`
		AssetIdQuote string  `json:"asset_id_quote"`
		Rate         float64 `json:"rate"`
	}

	Error struct {
		Message string `json:"error"`
	}

	Config struct {
		ApiKey   string `config:"COIN_API_KEY,required"`
		Endpoint string `config:"COIN_ENDPOINT"`
	}
)

type Client interface {
	ExchangeRate(ctx context.Context, assetBase string, assetQuote string) (float64, error)
}

type client struct {
	httpClient *http.Client
	cfg        *Config
}

const (
	defaultEndpoint = "https://rest.coinapi.io"
)

func NewClient() Client {
	cfg := Config{
		// ApiKey
		Endpoint: defaultEndpoint,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := confita.NewLoader(env.NewBackend()).Load(ctx, &cfg); err != nil {
		log.Fatalf("sendsay: failed to load config: %s", err)
	}

	return &client{
		cfg:        &cfg,
		httpClient: http.DefaultClient,
	}
}
