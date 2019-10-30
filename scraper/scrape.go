package scraper

import (
	"log"
	"strings"

	"github.com/gocolly/colly"
)

// Scrape is responsible for main scraping logic
func (s *Scraper) Scrape(scrapedEmails *[]string) {
	c := colly.NewCollector(
		colly.MaxDepth(s.MaxDepth),                               // Control the link following
		colly.Async(s.Async),                                     // Execute concurrently
		colly.AllowedDomains(prepareAllowedDomain(s.Website)...), // Ignore 3rd party links
	)

	if s.Recursively {
		// Find and visit all links
		c.OnHTML("a", func(e *colly.HTMLElement) {
			if err := e.Request.Visit(e.Attr("href")); err != nil {
				// Ignore already visited error, this appears too often
				if err != colly.ErrAlreadyVisited {
					if s.PrintLogs {
						log.Println("error while linking: ", err.Error())
					}
				}
			}
		})
	}

	// Parse emails on each downloaded page
	c.OnScraped(func(response *colly.Response) {
		parseEmails(response.Body, scrapedEmails)
	})

	// Start the scrape
	if err := c.Visit(s.Website); err != nil {
		if s.PrintLogs {
			log.Println("error while visiting: ", err.Error())
		}
	}

	// Wait for concurrent scrapes to finish
	c.Wait()
}

// Trim the input domain to whitelist root
func prepareAllowedDomain(domain string) []string {
	domain = strings.TrimLeft(domain, "http://")
	domain = strings.TrimLeft(domain, "https://")
	domain = strings.TrimLeft(domain, "wwww.")
	return []string{domain, "www." + domain, "*." + domain}
}
