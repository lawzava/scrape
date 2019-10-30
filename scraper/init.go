package scraper

// Goal of the scrape
type Scraper struct {
	Website string
	Parameters
}

// Configuration for the scrape
type Parameters struct {
	Emails      bool
	Recursively bool
	Async       bool
	MaxDepth    int
	PrintLogs   bool
}

// Initiate new scraper
func New(website string, parameters Parameters) *Scraper {
	var s Scraper

	s.Website = website
	s.Parameters = parameters

	return &s
}
