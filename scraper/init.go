package scraper

import "log"

// Goal of the scrape
type Scraper struct {
	Parameters
}

// Configuration for the scrape
type Parameters struct {
	Website             string
	MaxDepth            int
	Emails              bool
	Recursively         bool
	Async               bool
	PrintLogs           bool
	FollowExternalLinks bool
	JSWait              bool
}

// Initiate new scraper
func New(parameters Parameters) *Scraper {
	return &Scraper{
		Parameters: parameters,
	}
}

func (s *Scraper) Log(v ...interface{}) {
	if s.PrintLogs {
		log.Println(v...)
	}
}

func (s *Scraper) GetWebsite(secure bool) string {
	if secure {
		return "https://" + s.Website
	}

	return "http://" + s.Website
}
