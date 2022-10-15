package main

import (
	"github.com/lawzava/emailscraper"
)

//nolint:gochecknoglobals // allow global var here
var (
	scraperParameters emailscraper.Config
	url               string
	output            string
	outputWithURL     bool
)

//nolint:gochecknoinits // required by github.com/spf13/cobra
func init() {
	rootCmd.PersistentFlags().StringVarP(&url,
		"website", "w", "https://lawzava.com", "Website to scrape")
	rootCmd.PersistentFlags().BoolVar(&scraperParameters.Recursively,
		"recursively", true, "Scrape website recursively")
	rootCmd.PersistentFlags().IntVarP(&scraperParameters.MaxDepth,
		"depth", "d", 3, "Max depth to follow when scraping recursively") //nolint:gomnd // allow default max depth
	rootCmd.PersistentFlags().BoolVar(&scraperParameters.Async,
		"async", true, "Scrape website pages asynchronously")
	rootCmd.PersistentFlags().BoolVar(&scraperParameters.Debug,
		"debug", false, "Print debug logs")
	rootCmd.PersistentFlags().BoolVar(&scraperParameters.FollowExternalLinks,
		"follow-external", false, "Follow external 3rd party links within website")
	rootCmd.PersistentFlags().BoolVar(&scraperParameters.EnableJavascript,
		"js", false, "Enables EnableJavascript execution await")
	rootCmd.PersistentFlags().IntVar(&scraperParameters.Timeout,
		"timeout", 0, "If > 0, specify a timeout (seconds) for js execution await")
	rootCmd.PersistentFlags().StringVar(&output,
		"output", outputPlain, "Output type to use (default 'plain', supported: 'csv', 'json')")
	rootCmd.PersistentFlags().BoolVar(&outputWithURL,
		"output-with-url", false, "Adds URL to output with each email")
}
