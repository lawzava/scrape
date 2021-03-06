package main

import (
	"fmt"
	"log"
	"os"

	"github.com/lawzava/emailscraper"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// nolint:gochecknoglobals // allow global var here
var (
	scraperParameters emailscraper.Config
	url               string
)

// nolint:exhaustivestruct,gochecknoglobals // not valid requirement for this use case
var rootCmd = &cobra.Command{
	Use:   "scrape",
	Short: "CLI utility to scrape emails from websites",
	Long:  `CLI utility that scrapes emails from specified website recursively and concurrently`,
	Run: func(cmd *cobra.Command, args []string) {
		scraper := emailscraper.New(scraperParameters)

		// Scrape for emails
		scrapedEmails, err := scraper.Scrape(url)
		if err != nil {
			log.Fatal(err)
		}

		for _, email := range scrapedEmails {
			fmt.Println(email) // nolint:forbidigo // allow println here for non intrusive response
		}
	},
}

// nolint:gochecknoinits // required by github.com/spf13/cobra
func init() {
	rootCmd.PersistentFlags().StringVarP(&url,
		"website", "w", "https://lawzava.com", "Website to scrape")
	rootCmd.PersistentFlags().BoolVar(&scraperParameters.Recursively,
		"recursively", true, "Scrape website recursively")
	rootCmd.PersistentFlags().IntVarP(&scraperParameters.MaxDepth,
		"depth", "d", 3, "Max depth to follow when scraping recursively")
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
}
