package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/lawzava/scrape/scraper"

	"github.com/spf13/cobra"
)

// nolint:gochecknoglobals // allow global var here
var scraperParameters scraper.Parameters

// nolint:exhaustivestruct,gochecknoglobals // not valid requirement for this use case
var rootCmd = &cobra.Command{
	Use:   "scrape",
	Short: "CLI utility to scrape emails from websites",
	Long:  `CLI utility that scrapes emails from specified website recursively and concurrently`,
	Run: func(cmd *cobra.Command, args []string) {
		scraper := scraper.New(scraperParameters)

		// Scrape for emails
		scrapedEmails, err := scraper.Scrape()
		if err != nil {
			log.Fatal(err)
		}

		for _, email := range scrapedEmails {
			fmt.Println(email) // nolint:forbidigo // allow println here for non intrusive response
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

// nolint:gochecknoinits // required by github.com/spf13/cobra
func init() {
	rootCmd.PersistentFlags().StringVarP(&scraperParameters.Website,
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
	rootCmd.PersistentFlags().BoolVar(&scraperParameters.Emails,
		"emails", true, "Scrape emails")
	rootCmd.PersistentFlags().BoolVar(&scraperParameters.JS,
		"js", false, "Enables JS execution await")
	rootCmd.PersistentFlags().IntVar(&scraperParameters.Timeout,
		"timeout", 0, "If > 0, specify a timeout (seconds) for js execution await")
}
