package main

import (
	"context"

	"go.uber.org/fx"

	"github.com/igilgyrg/statbot/api/clients/coin"
	"github.com/igilgyrg/statbot/internal/currency"
	"github.com/igilgyrg/statbot/internal/storage/binary"
)

const filename = "store.json"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client := coin.NewClient()

	repository := binary.NewRepository(filename)

	service := currency.NewService(repository, client)

	application := fx.New(
		fx.Supply(ctx, client, repository, service),
	)

	application.Run()
}
