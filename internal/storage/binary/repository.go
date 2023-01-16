package binary

import (
	"github.com/igilgyrg/statbot/internal/storage"
)

type repository struct {
	filename string
}

func NewRepository(filename string) storage.Storage {
	return &repository{filename: filename}
}
