package currency

import (
	"context"

	"github.com/igilgyrg/statbot/api/clients"
	"github.com/igilgyrg/statbot/internal/storage"
	"github.com/igilgyrg/statbot/model"
)

type Service interface {
	Report(ctx context.Context) (model.Capital, error)
	Save(ctx context.Context, save *model.DepositSave) error
}

type service struct {
	storage storage.Storage
	client  clients.Client
}

func NewService(storage storage.Storage, client clients.Client) Service {
	return &service{storage: storage, client: client}
}
