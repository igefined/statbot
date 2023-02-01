package supabase

import (
	"context"
	"strings"

	"github.com/igilgyrg/statbot/internal/storage"
)

func (r repository) Save(ctx context.Context, asset storage.Deposit) error {
	q := "insert into coins_stats(symbol, count, purchase_price) values($1, $2, $3) returning id"
	exec := r.qb.Querier().QueryRow(ctx, q, strings.ToUpper(asset.Symbol), asset.Count, asset.PurchasePrice)

	return exec.Scan(&asset.Id)
}
