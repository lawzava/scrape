package scraper

import "log"

// Goal of the scrape
type Scraper struct {
	Parameters
}

// Configuration for the scrape
type Parameters struct {
	Website             string
	Emails              bool
	Recursively         bool
	Async               bool
	MaxDepth            int
	PrintLogs           bool
	FollowExternalLinks bool
}

// Initiate new scraper
func New(parameters Parameters) *Scraper {
	var s Scraper
	s.Parameters = parameters
	return &s
}

func (s *Scraper) Log(v ...interface{}) {
	if s.PrintLogs {
		log.Println(v)
	}
}
