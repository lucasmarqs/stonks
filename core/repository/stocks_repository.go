package repository

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lucasmarqs/stonks/core/model"
)

type StocksRepository struct {
	DB *sqlx.DB
}

func (repo StocksRepository) Find(name string) (model.Stock, error) {
	query := `SELECT * FROM stocks WHERE name = ?`

	stock := model.Stock{}
	err := repo.DB.Get(&stock, query, name)

	return stock, err
}

func (repo StocksRepository) FindOrCreate(name string) (model.Stock, error) {
	stock, err := repo.Find(name)
	if err == nil {
		return stock, nil
	} else if err.Error() != sql.ErrNoRows.Error() {
		return stock, err
	}

	insert := `INSERT INTO stocks (name) values (?)`
	if _, err := repo.DB.Exec(insert, name); err != nil {
		return stock, err
	}

	return repo.Find(name)
}

func (repo StocksRepository) SaveNewQuote(stock *model.Stock, quote float64) error {
	now := time.Now()
	stock.LastQuote = &quote
	stock.LastQuoteDate = &now

	query := `UPDATE stocks SET last_quote = ?, last_quote_date = ? WHERE id = ?`
	_, err := repo.DB.Exec(query, quote, now, stock.ID)
	return err
}
