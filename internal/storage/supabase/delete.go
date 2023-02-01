package supabase

import (
	"context"
)

func (r repository) Delete(ctx context.Context, id int64) (err error) {
	_, err = r.qb.Querier().Query(ctx, "delete from coins_stats where id = $1", id)

	return
}

func (r repository) DeleteAllByIds(ctx context.Context, ids []int64) (err error) {
	_, err = r.qb.Querier().Exec(ctx, "delete from coins_stats where id = any($1)", ids)

	return
}
