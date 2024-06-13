package domain

import "time"

type ConvertingInfo struct {
	From  string
	To    string
	Count int64 // в копейках
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
