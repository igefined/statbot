package model

import "time"

type USDT struct {
	Asset string
	Rate  float64
	Time  time.Time
}
