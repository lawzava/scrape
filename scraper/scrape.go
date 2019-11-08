package scraper

import (
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

// Scrape is responsible for main scraping logic
func (s *Scraper) Scrape(scrapedEmails *[]string) error {
	// Configure colly
	c := colly.NewCollector()

	c.MaxDepth = s.MaxDepth
	c.Async = s.Async
	if !s.FollowExternalLinks {
		allowedDomains, err := prepareAllowedDomain(s.Website)
		if err != nil {
			return err
		}
		c.AllowedDomains = allowedDomains
	}
	s.Website = trimProtocol(s.Website)

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

	// Wait for concurrent scrapes to finish
	c.Wait()

	if scrapedEmails == nil {
		// Start the scrape
		if err := c.Visit(s.GetWebsite(false)); err != nil {
			s.Log("error while visiting: ", err.Error())
		}

		// Wait for concurrent scrapes to finish
		c.Wait()
	}

	return nil
}

// Trim the input domain to whitelist root
func prepareAllowedDomain(requestURL string) ([]string, error) {
	requestURL = "https://" + trimProtocol(requestURL)
	u, err := url.ParseRequestURI(requestURL)
	if err != nil {
		return nil, err
	}
	hostname := u.Hostname()
	domain := strings.TrimLeft(hostname, "wwww.")
	return []string{
		domain,
		"www." + domain,
		"http://" + domain,
		"https://" + domain,
		"http://www." + domain,
		"https://www." + domain,
	}, nil
}

func trimProtocol(requestURL string) string {
	return strings.Trim(strings.Trim(requestURL, "http://"), "https://")
}
