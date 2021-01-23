package scraper

import (
	"net/url"
	"strings"
)

// Trim the input domain to whitelist root
func prepareAllowedDomain(requestURL string) ([]string, error) {
	requestURL = "https://" + trimProtocol(requestURL)

	u, err := url.ParseRequestURI(requestURL)
	if err != nil {
		return nil, err
	}

	domain := strings.TrimPrefix(u.Hostname(), "www.")

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
