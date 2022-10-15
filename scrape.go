package main

import (
	"log"

	"github.com/lawzava/emailscraper"
	"github.com/spf13/cobra"
)

//nolint:exhaustivestruct,gochecknoglobals,exhaustruct // not valid requirement for this use case
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

		handleOutput(scrapedEmails)
	},
}
