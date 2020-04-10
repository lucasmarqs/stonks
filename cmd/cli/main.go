package main

import (
	"fmt"

	"github.com/lucasmarqs/stonks/core/scraping/yahoo"
)

func main() {
	scrapper := yahoo.NewYahoo()

	names := []string{"nslu11", "itsa4", "vlol11"}

	for _, n := range names {
		fmt.Printf("%.2f\n", scrapper.QuotationFor(n))
	}
}
