package scraper

import (
	"log"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

// Scrape is responsible for main scraping logic
func (s *Scraper) Scrape(scrapedEmails *[]string) error {
	allowedDomains, err := prepareAllowedDomain(s.Website)
	if err != nil {
		return err
	}
	c := colly.NewCollector(
		colly.MaxDepth(s.MaxDepth),              // Control the link following
		colly.Async(s.Async),                    // Execute concurrently
		colly.AllowedDomains(allowedDomains...), // Ignore 3rd party links
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
			log.Printf("error while visiting: %s\n", err.Error())
		}
	}

	// Wait for concurrent scrapes to finish
	c.Wait()

	return nil
}

// Trim the input domain to whitelist root
func prepareAllowedDomain(requestURL string) ([]string, error) {
	u, err := url.ParseRequestURI(requestURL)
	if err != nil {
		return nil, err
	}
	hostname := u.Hostname()
	domain := strings.TrimLeft(hostname, "wwww.")
	return []string{
		domain,
		"www." + domain,
		"*." + domain,
		"http://" + domain,
		"https://" + domain,
		"http://www." + domain,
		"https://www." + domain,
	}, nil
}
