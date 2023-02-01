package supabase

import (
	"context"
	"github.com/igilgyrg/statbot/internal/storage"
)

func (r repository) Update(ctx context.Context, asset *storage.Deposit) (err error) {
	q := `update coins_stats set count = $2 where id = $1`
	_, err = r.qb.Querier().Query(ctx, q, asset.Id, asset.Count)

	return
}
