package currency

import (
	"context"
	"fmt"

	"github.com/igilgyrg/statbot/internal/storage"
	"github.com/igilgyrg/statbot/model"
	"github.com/igilgyrg/statbot/utils"
)

const usdt = "USDT"

func (s service) Report(ctx context.Context) (capital model.Capital, err error) {
	deposits, err := s.storage.List(ctx)
	if err != nil {
		err = fmt.Errorf("error list storage: %w", err)

		return
	}

	depositOnMap := toMap(deposits)
	capital = s.GetInfoAndAverage(ctx, depositOnMap)

	return
}

func (s service) GetInfoAndAverage(ctx context.Context, deposits map[string][]storage.Deposit) (result model.Capital) {
	result = make(map[string]model.Deposit, len(deposits))

	for k, v := range deposits {
		var (
			pcp, pv, vn float64
		)

		rate, _ := s.client.ExchangeRate(ctx, k, usdt)

		for _, d := range v {
			pv += d.Count * d.PurchasePrice
			vn += d.Count
		}
		pcp = pv / vn

		result[k] = model.Deposit{
			Count:         utils.RoundFloat2Decimal(vn),
			PurchasePrice: utils.RoundFloat2Decimal(pcp),
			ActualPrice:   utils.RoundFloat2Decimal(rate),
			Difference:    utils.RoundFloat2Decimal(differenceInPercent(pcp, rate)),
		}
	}

	return
}

func toMap(deposits []storage.Deposit) (result map[string][]storage.Deposit) {
	result = make(map[string][]storage.Deposit, len(deposits))

	for i := range deposits {
		deposit := deposits[i]
		result[deposit.Symbol] = append(result[deposit.Symbol], deposit)
	}

	return
}

func differenceInPercent(price, rate float64) float64 {
	return (rate - price) / price * 100
}
