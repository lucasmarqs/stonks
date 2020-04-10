package scraping

type Scraper interface {
	QuotationFor(string) float64
}
