package coin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *client) ExchangeRate(ctx context.Context, assetBase string, assetQuote string) (rates float64, err error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/v1/exchangerate/%s/%s", c.cfg.Endpoint, assetBase, assetQuote), nil)
	if err != nil {
		err = fmt.Errorf("coin exchange rate make request: %w", err)

		return
	}

	req.Header.Set("X-CoinAPI-Key", c.cfg.ApiKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		err = fmt.Errorf("coin exchange rate do request: %w", err)

		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		var errResponse Error

		if err = json.NewDecoder(resp.Body).Decode(&errResponse); err != nil {
			err = fmt.Errorf("coin exchange rate error parse object: %w", err)

			return
		}

		err = fmt.Errorf("coin exchange rate error request: %s", errResponse.Message)

		return
	}

	var ratesResponse Rates

	if err = json.NewDecoder(resp.Body).Decode(&ratesResponse); err != nil {
		err = fmt.Errorf("coin exchange rate error parse object: %w", err)

		return
	}

	rates = ratesResponse.Rate

	return
}
