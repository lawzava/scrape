package scraper

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/chromedp/chromedp"

	"github.com/gocolly/colly"
)

// Scrape is responsible for main scraping logic
func (s *Scraper) Scrape(scrapedEmails *[]string) error {
	// Initiate colly
	c := colly.NewCollector()

	c.MaxDepth = s.MaxDepth
	c.Async = s.Async
	s.Website = trimProtocol(s.Website)

	if !s.FollowExternalLinks {
		allowedDomains, err := prepareAllowedDomain(s.Website)
		if err != nil {
			return err
		}

		c.AllowedDomains = allowedDomains
	}

	if s.JSWait {
		c.OnResponse(func(response *colly.Response) {
			if err := initiateChromeSession(response); err != nil {
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

	if scrapedEmails == nil {
		// Start the scrape on insecure url
		if err := c.Visit(s.GetWebsite(false)); err != nil {
			s.Log("error while visiting: ", err.Error())
		}

		c.Wait() // Wait for concurrent scrapes to finish
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

	domain := strings.TrimPrefix(u.Hostname(), "wwww.")

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
	return strings.TrimPrefix(strings.TrimPrefix(requestURL, "http://"), "https://")
}

func initiateChromeSession(response *colly.Response) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string
	if err := chromedp.Run(ctx, chromedp.Navigate(response.Request.URL.String()),
		chromedp.InnerHTML("html", &res), // Scrape whole rendered page
	); err != nil {
		return fmt.Errorf("executing chromedp: %w", err)
	}

	response.Body = []byte(res)

	return nil
}
