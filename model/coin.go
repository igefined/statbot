package model

import (
	"errors"
	"time"
)

type (
	USDT struct {
		Asset string
		Rate  float64
		Time  time.Time
	}

	Deposit struct {
		Count         float64 `json:"count"`
		PurchasePrice float64 `json:"purchase_price"`
		ActualPrice   float64 `json:"actual_price"`
		Difference    float64 `json:"difference"` // percent
	}

	DepositSave struct {
		Symbol        string
		Count         float64
		PurchasePrice float64
	}

	Capital map[string]Deposit
)

var ErrCoinNotFound = errors.New("currency not founded")
