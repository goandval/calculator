package domain

import (
	"time"
)

type RateRequest struct {
	From string
	To   string
}

type RateResponse struct {
	Count   int64
	Updated time.Time
}

type Response struct {
	From        Currency
	To          Currency
	LastUpdated time.Time
}

type Currency struct {
	Name  string
	Count int64
}

// EUR, USD, CNY, USDT, USDC, ETH
