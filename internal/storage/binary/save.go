package binary

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/igilgyrg/statbot/internal/storage"
)

func (r repository) Save(ctx context.Context, asset storage.Deposit) error {
	capital, err := r.List(ctx)
	if err != nil {
		return err
	}

	capital = append(capital, asset)

	bytes, err := json.Marshal(capital)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(r.filename, bytes, 0644)
}
