package storage

import "context"

type Deposit struct {
	Symbol        string  `json:"symbol"`
	Count         float64 `json:"count"`
	PurchasePrice float64 `json:"purchase_price"`
}

type Storage interface {
	List(ctx context.Context) ([]Deposit, error)
	Save(ctx context.Context, asset Deposit) error
}
