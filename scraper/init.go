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
	Recursively         bool
	Async               bool
	JS                  bool
	FollowExternalLinks bool
	Debug               bool

	// Objects to scrape
	Emails bool
}

// Initiate new scraper
func New(parameters Parameters) *Scraper {
	return &Scraper{
		Parameters: parameters,
	}
}

func (s *Scraper) Log(v ...interface{}) {
	if s.Debug {
		log.Println(v...)
	}
}

func (s *Scraper) GetWebsite(secure bool) string {
	if secure {
		return "https://" + s.Website
	}

	return "http://" + s.Website
}
