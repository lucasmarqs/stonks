package yahoo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

const markup = "div#quote-header-info div:nth-child(3) span:nth-child(1)"

var priceRegexp = regexp.MustCompile(`\d+\.\d{2}`)

type Yahoo struct {
	collector *colly.Collector
}

func NewYahoo() *Yahoo {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US")
	})

	return &Yahoo{collector: c}
}

func (y Yahoo) QuotationFor(name string) (quotation float64) {
	y.collector.OnHTML(markup, func(e *colly.HTMLElement) {
		if ok := priceRegexp.MatchString(e.Text); !ok {
			return
		}

		val, err := strconv.ParseFloat(e.Text, 32)
		if err != nil {
			return
		}
		quotation = val
	})

	param := fmt.Sprintf("%s.SA", strings.ToUpper(name))
	y.collector.Visit(fmt.Sprintf("https://finance.yahoo.com/quote/%s", param))

	return
}
