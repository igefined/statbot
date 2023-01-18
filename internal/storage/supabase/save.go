package supabase

import (
	"context"

	"github.com/igilgyrg/statbot/internal/storage"
)

func (r repository) Save(ctx context.Context, asset storage.Deposit) (err error) {
	q := "insert into coins_stats(symbol, count, purchas_price) values($1, $2, $3) returning id"
	exec := r.qb.Querier().QueryRow(ctx, q, asset.Symbol, asset.Count, asset.PurchasePrice)

	return exec.Scan(&asset.Id)
}
