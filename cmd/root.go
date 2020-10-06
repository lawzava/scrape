package cmd

import (
	"fmt"
	"lawzava/scrape/scraper"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var scraperParameters scraper.Parameters

var rootCmd = &cobra.Command{
	Use:   "scrape",
	Short: "CLI utility to scrape emails from websites",
	Long:  `CLI utility that scrapes emails from specified website recursively and concurrently`,
	Run: func(cmd *cobra.Command, args []string) {
		scrap := scraper.New(scraperParameters)

		// Scrape for emails
		var scrapedEmails []string
		if err := scrap.Scrape(&scrapedEmails); err != nil {
			log.Fatal(err)
		}

		for _, email := range scrapedEmails {
			fmt.Println(email)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

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
