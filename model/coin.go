package model

import (
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

	Capital map[string]Deposit
)
