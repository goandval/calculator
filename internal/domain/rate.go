package domain

import (
	"errors"
	"fmt"
	"sync/atomic"
	"time"
)

type RateMap map[string]map[string]int64

var (
	errUpdated      = errors.New("updated time is invalid, data may be inconsistent")
	errDataNotFound = errors.New("currency data not found")
)

type ExchangeRate struct {
	data    atomic.Pointer[RateMap]
	updated atomic.Value
}

func (r *ExchangeRate) Rate(req RateRequest) (RateResponse, error) {
	const op = "domain.Rate"
	rate := r.data.Load()
	if rate == nil {
		return RateResponse{}, fmt.Errorf("%s: %w", op, errDataNotFound)
	}
	updated := r.updated.Load()
	if updated == nil {
		return RateResponse{}, fmt.Errorf("%s: %w", op, errUpdated)
	}
	updatedTime, ok := updated.(time.Time)
	if !ok {
		return RateResponse{}, fmt.Errorf("%s: %w", op, errUpdated)
	}
	return RateResponse{
		Count:   int64((*rate)[req.From][req.To]),
		Updated: updatedTime,
	}, nil
}

func (r *ExchangeRate) Update(rate RateMap) {

}
