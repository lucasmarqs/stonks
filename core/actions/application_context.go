package actions

import (
	"github.com/jmoiron/sqlx"
	"github.com/lucasmarqs/stonks/core/database"
	"github.com/lucasmarqs/stonks/core/repository"
	"github.com/lucasmarqs/stonks/core/scraping"
	"github.com/lucasmarqs/stonks/core/scraping/yahoo"
)

type ApplicationContext struct {
	db               *sqlx.DB
	stocksRepository repository.StocksRepository
	scrapper         scraping.Scraper
}

func StartApp() ApplicationContext {
	db := database.EstablishConnection()
	return ApplicationContext{
		db:               db,
		stocksRepository: repository.StocksRepository{DB: db},
		scrapper:         yahoo.NewYahoo(),
	}
}

func (app ApplicationContext) Shutdown() {
	app.db.Close()
}
