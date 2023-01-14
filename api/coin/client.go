package coin

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Config struct {
	ApiKey   string `config:"COIN_API_KEY,required"`
	Endpoint string `config:"COIN_ENDPOINT"`
}

type Client interface {
	ExchangeRate(ctx context.Context, assetBase string, assetQuote string) error
}

type client struct {
	httpClient *http.Client
	cfg        *Config
}

const (
	defaultEndpoint = "https://rest.coinapi.io"
)

func New() Client {
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
