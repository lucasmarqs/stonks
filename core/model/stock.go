package model

import "time"

type Stock struct {
	ID            uint       `db:"id"`
	Name          string     `db:"name"`
	LastQuote     *float64   `db:"last_quote"`
	LastQuoteDate *time.Time `db:"last_quote_date"`
}
