/*
Copyright © 2019 Law Zava <i@lawzava.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/lawzava/scrape/scraper"

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
	rootCmd.PersistentFlags().StringVarP(&scraperParameters.Website, "website", "w", "https://lawzava.com", "Website to scrape")
	rootCmd.PersistentFlags().BoolVar(&scraperParameters.Recursively, "recursively", true, "Scrape website recursively")
	rootCmd.PersistentFlags().IntVarP(&scraperParameters.MaxDepth, "depth", "d", 3, "Max depth to follow when scraping recursively")
	rootCmd.PersistentFlags().BoolVar(&scraperParameters.Async, "async", true, "Scrape website pages asynchronously")
	rootCmd.PersistentFlags().BoolVar(&scraperParameters.PrintLogs, "logs", false, "Print debug logs")
	rootCmd.PersistentFlags().BoolVar(&scraperParameters.FollowExternalLinks, "follow-external", false, "Follow external 3rd party links within website")
	rootCmd.PersistentFlags().BoolVar(&scraperParameters.Emails, "emails", true, "Scrape emails")
	rootCmd.PersistentFlags().BoolVar(&scraperParameters.JSWait, "js-wait", false, "Should wait for JS to execute")
}
