package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/lawzava/scrape/scraper"
)

func main() {

	// Parse flags
	websiteToScrape := flag.String("website", "https://v0.vc", "specify a website to scrape")
	recursively := flag.Bool("recursively", true, "scan website recursively")
	async := flag.Bool("async", true, "scan website concurrently")
	maxDepth := flag.Int("depth", 1, "maximum depth for recursive scan")
	printLogs := flag.Bool("log", false, "print logs")
	flag.Parse()

	// Initiate scraper
	scrap := scraper.New(*websiteToScrape, scraper.Parameters{
		Emails:      true,
		Recursively: *recursively,
		Async:       *async,
		MaxDepth:    *maxDepth,
		PrintLogs:   *printLogs,
	})

	// Scrape for emails
	var scrapedEmails []string
	if err := scrap.Scrape(&scrapedEmails); err != nil {
		log.Fatal(err)
	}

	if *printLogs {
		fmt.Printf("\n\n\n")
		fmt.Println("=================================")
		fmt.Println("Scrape finished. Results:")
		fmt.Println(" ")
	}
	for _, email := range scrapedEmails {
		fmt.Println(email)
	}
}
