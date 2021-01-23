package scraper

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

// Scrape is responsible for main scraping logic
func (s *Scraper) Scrape(scrapedEmails *[]string) error {
	// Initiate colly
	c := colly.NewCollector()

	c.Async = s.Async
	c.MaxDepth = s.MaxDepth
	s.Website = trimProtocol(s.Website)

	if !s.FollowExternalLinks {
		allowedDomains, err := prepareAllowedDomain(s.Website)
		if err != nil {
			return err
		}

		c.AllowedDomains = allowedDomains
	}

	if s.Debug {
		c.SetDebugger(&debug.LogDebugger{})
	}

	if s.JS {
		c.OnResponse(func(response *colly.Response) {
			if err := initiateScrapingFromChrome(response, s.Timeout); err != nil {
				s.Log(err)
				return
			}
		})
	}

	if s.Recursively {
		// Find and visit all links
		c.OnHTML("a", func(e *colly.HTMLElement) {
			s.Log("visiting: ", e.Attr("href"))
			if err := e.Request.Visit(e.Attr("href")); err != nil {
				// Ignore already visited error, this appears too often
				if err != colly.ErrAlreadyVisited {
					s.Log("error while linking: ", err.Error())
				}
			}
		})
	}

	// Parse emails on each downloaded page
	c.OnScraped(func(response *colly.Response) {
		parseEmails(response.Body, scrapedEmails)
	})

	// Start the scrape
	if err := c.Visit(s.GetWebsite(true)); err != nil {
		s.Log("error while visiting: ", err.Error())
	}

	c.Wait() // Wait for concurrent scrapes to finish

	if scrapedEmails == nil || len(*scrapedEmails) == 0 {
		// Start the scrape on insecure url
		if err := c.Visit(s.GetWebsite(false)); err != nil {
			s.Log("error while visiting: ", err.Error())
		}

		c.Wait() // Wait for concurrent scrapes to finish
	}

	return nil
}
