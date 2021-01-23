package scraper

import (
	"errors"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

// Scrape is responsible for main scraping logic.
func (s *Scraper) Scrape() ([]string, error) {
	// Initiate colly
	c := colly.NewCollector()

	c.Async = s.Async
	c.MaxDepth = s.MaxDepth
	s.Website = trimProtocol(s.Website)
	e := emails{}

	if !s.FollowExternalLinks {
		allowedDomains, err := prepareAllowedDomain(s.Website)
		if err != nil {
			return nil, err
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
		c.OnHTML("a[href]", func(el *colly.HTMLElement) {
			s.Log("visiting: ", el.Attr("href"))
			if err := el.Request.Visit(el.Attr("href")); err != nil {
				// Ignore already visited error, this appears too often
				if !errors.Is(err, colly.ErrAlreadyVisited) {
					s.Log("error while linking: ", err.Error())
				}
			}
		})
	}

	// Parse emails on each downloaded page
	c.OnScraped(func(response *colly.Response) {
		e.parseEmails(response.Body)
	})

	// cloudflare encoded email support
	c.OnHTML("span[data-cfemail]", func(el *colly.HTMLElement) {
		e.parseCloudflareEmail(el.Attr("data-cfemail"))
	})

	// Start the scrape
	if err := c.Visit(s.GetWebsite(true)); err != nil {
		s.Log("error while visiting: ", err.Error())
	}

	c.Wait() // Wait for concurrent scrapes to finish

	if e.emails == nil || len(e.emails) == 0 {
		// Start the scrape on insecure url
		if err := c.Visit(s.GetWebsite(false)); err != nil {
			s.Log("error while visiting: ", err.Error())
		}

		c.Wait() // Wait for concurrent scrapes to finish
	}

	return e.emails, nil
}
