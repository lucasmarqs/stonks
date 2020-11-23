package actions

import (
	"github.com/lucasmarqs/stonks/core/model"
)

func (app ApplicationContext) QuoteStock(name string) (model.Stock, error) {
	stock, err := app.stocksRepository.FindOrCreate(name)
	if err != nil {
		return stock, err
	}

	quote := app.scrapper.QuotationFor(stock.Name)
	if err := app.stocksRepository.SaveNewQuote(&stock, quote); err != nil {
		return model.Stock{}, err
	}

	return stock, nil
}
