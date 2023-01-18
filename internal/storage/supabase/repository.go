package supabase

import (
	"github.com/igilgyrg/statbot/internal/storage"
	"github.com/igilgyrg/statbot/schema"
)

type repository struct {
	qb *schema.QBuilder
}

func NewRepository(builder *schema.QBuilder) storage.Storage {
	return &repository{qb: builder}
}
