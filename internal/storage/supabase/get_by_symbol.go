package supabase

import (
	"context"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/igilgyrg/statbot/internal/storage"
	"strings"
)

func (r repository) GetBySymbol(ctx context.Context, symbol string) (deposit []storage.Deposit, err error) {
	q := `select * from coins_stats cs where upper(cs.symbol) = upper($1)`
	err = pgxscan.Select(ctx, r.qb.Querier(), &deposit, q, strings.ToUpper(symbol))

	return
}
