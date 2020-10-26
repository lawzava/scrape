package scraper

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"net/url"
	"strings"
	"time"
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

func initiateScrapingFromChrome(response *colly.Response, timeout int) error {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3830.0 Safari/537.36"), // nolint
		chromedp.WindowSize(1920, 1080),
		chromedp.NoFirstRun,
		chromedp.Headless,
		chromedp.DisableGPU,
	}

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	if timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
		defer cancel()
	}

	var res string
	if err := chromedp.Run(ctx, chromedp.Navigate(response.Request.URL.String()),
		chromedp.InnerHTML("html", &res), // Scrape whole rendered page
	); err != nil {
		return fmt.Errorf("executing chromedp: %w", err)
	}
	response.Body = []byte(res)

	return nil
}
