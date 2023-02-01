package storage

import "context"

type Deposit struct {
	Id            int64   `json:"id" db:"id"`
	Symbol        string  `json:"symbol" db:"symbol"`
	Count         float64 `json:"count" db:"count"`
	PurchasePrice float64 `json:"purchase_price" db:"purchase_price"`
}

type Storage interface {
	List(ctx context.Context) ([]Deposit, error)
	Save(ctx context.Context, asset Deposit) error
	GetBySymbol(ctx context.Context, symbol string) ([]Deposit, error)
	Delete(ctx context.Context, id int64) error
	DeleteAllByIds(ctx context.Context, ids []int64) error
	Update(ctx context.Context, asset *Deposit) error
}
