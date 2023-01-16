package currency

import (
	"context"
	"strings"

	"github.com/igilgyrg/statbot/internal/storage"
	"github.com/igilgyrg/statbot/model"
)

func (s service) Save(ctx context.Context, save *model.DepositSave) error {
	save.Symbol = strings.ToUpper(save.Symbol)
	rate, err := s.client.ExchangeRate(ctx, save.Symbol, "USDT")
	if err != nil || rate == 0.0 {
		return model.ErrCoinNotFound
	}

	return s.storage.Save(ctx, storage.Deposit{Symbol: save.Symbol, Count: save.Count, PurchasePrice: save.PurchasePrice})
}
