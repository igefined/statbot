package supabase

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/igilgyrg/statbot/internal/storage"
)

func (r repository) List(ctx context.Context) (list []storage.Deposit, err error) {
	err = pgxscan.Select(ctx, r.qb.Querier(), &list, "select * from coins_stats order by count desc")

	return
}
