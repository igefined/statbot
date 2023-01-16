package binary

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/igilgyrg/statbot/internal/storage"
)

func (r repository) List(ctx context.Context) (deposits []storage.Deposit, err error) {
	jsonFile, err := os.OpenFile("store.json", os.O_CREATE, os.ModePerm)
	if err != nil {
		err = fmt.Errorf("error open store file binary repository: %w", err)

		return
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		err = fmt.Errorf("error save binary repository: %w", err)

		return
	}

	deposits = make([]storage.Deposit, 0)
	err = json.Unmarshal(bytes, &deposits)
	if err != nil {
		err = fmt.Errorf("error unmarshal capital from binary repository: %w", err)

		return
	}

	return
}
