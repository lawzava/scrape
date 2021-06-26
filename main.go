package main

import (
	"encoding/json"
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
	output            string
	outputWithURL     bool
)

const (
	outputPlain = "plain"
	outputCSV   = "csv"
	outputJSON  = "json"
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

		handleOutput(scrapedEmails)
	},
}

// nolint:gochecknoinits // required by github.com/spf13/cobra
func init() {
	rootCmd.PersistentFlags().StringVarP(&url,
		"website", "w", "https://lawzava.com", "Website to scrape")
	rootCmd.PersistentFlags().BoolVar(&scraperParameters.Recursively,
		"recursively", true, "Scrape website recursively")
	rootCmd.PersistentFlags().IntVarP(&scraperParameters.MaxDepth,
		"depth", "d", 3, "Max depth to follow when scraping recursively") // nolint:gomnd // allow default max depth
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

// nolint:forbidigo // allow println here for non intrusive response
func handleOutput(emails []string) {
	switch output {
	case outputCSV:
		for _, email := range emails {
			if outputWithURL {
				fmt.Print(url + ",")
			}

			fmt.Println(email)
		}
	case outputJSON:
		fmt.Println(prepareJSONOutput(emails))
	default:
		for _, email := range emails {
			if outputWithURL {
				fmt.Print(url + " ")
			}

			fmt.Println(email)
		}
	}
}

func prepareJSONOutput(emails []string) string {
	if outputWithURL {
		type outputFormat struct {
			URL   string `json:"url"`
			Email string `json:"email"`
		}

		out := make([]outputFormat, len(emails))
		for i, email := range emails {
			out[i].URL = url
			out[i].Email = email
		}

		b, err := json.Marshal(out)
		if err != nil {
			log.Fatal(fmt.Errorf("failed to marshal json with url response: %w", err))
		}

		return string(b)
	}

	type outputFormat struct {
		Email string `json:"email"`
	}

	out := make([]outputFormat, len(emails))
	for i, email := range emails {
		out[i].Email = email
	}

	b, err := json.Marshal(out)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to marshal json response: %w", err))
	}

	return string(b)
}
