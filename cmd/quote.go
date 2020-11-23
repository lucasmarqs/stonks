package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// quoteCmd represents the quote command
var (
	// flags
	ticker string

	// cmd
	quoteCmd = &cobra.Command{
		Use:   "quote",
		Short: "quote a stock",
		Run:   quote,
	}
)

func init() {
	rootCmd.AddCommand(quoteCmd)

	quoteCmd.Flags().StringVarP(&ticker, "ticker", "t", "", "Stock's ticker you want to quote")
	quoteCmd.MarkFlagRequired("ticker")
}

func quote(_ *cobra.Command, _ []string) {
	stock, err := app.QuoteStock(ticker)
	if err != nil {
		fmt.Printf("failed to get %s quotation. %s\n", ticker, err.Error())
	}

	fmt.Printf("%3.2f\n", *stock.LastQuote)
}
