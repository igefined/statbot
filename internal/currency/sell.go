package currency

import (
	"context"
	"github.com/igilgyrg/statbot/internal/storage"
	"github.com/igilgyrg/statbot/model"
)

func (s service) Sell(ctx context.Context, deposit *model.DepositPatch) (err error) {
	symbols, err := s.storage.GetBySymbol(ctx, deposit.Symbol)
	if err != nil {
		return
	}

	sumCount := 0.0
	averagePrice := 0.0
	ids := make([]int64, len(symbols))

	for i := range symbols {
		symbol := symbols[i]

		ids[i] = symbol.Id
		sumCount += symbol.Count
		averagePrice += symbol.PurchasePrice
	}

	if sumCount > deposit.Count {
		err = s.storage.DeleteAllByIds(ctx, ids)
		if err != nil {
			return
		}

		err = s.storage.Save(ctx, storage.Deposit{Symbol: deposit.Symbol, Count: sumCount, PurchasePrice: averagePrice})

		return
	}

	err = s.storage.DeleteAllByIds(ctx, ids)

	return
}
