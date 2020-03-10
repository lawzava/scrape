package scraper

import (
	"regexp"
	"strconv"
	"strings"

	"lawzava/scrape/tld"
)

// Initialize once
var reg = regexp.MustCompile(`([a-zA-Z0-9._-]+@([a-zA-Z0-9_-]+\.)+[a-zA-Z0-9_-]+)`)

// Parse any *@*.* string and append to the slice
func parseEmails(body []byte, scrapedEmails *[]string) {
	res := reg.FindAll(body, -1)

	for _, r := range res {
		email := string(r)
		if !isValidEmail(email) {
			continue
		}

		var skip bool
		// Check for already existing emails
		for _, existingEmail := range *scrapedEmails {
			if existingEmail == email {
				skip = true
				break
			}
		}

		if skip {
			continue
		}

		*scrapedEmails = append(*scrapedEmails, email)
	}
}

// Check if email looks valid
func isValidEmail(email string) bool {
	split := strings.Split(email, ".")
	if len(split) < 2 {
		return false
	}

	ending := split[len(split)-1]

	if len(ending) < 2 {
		return false
	}

	// check if TLD name actually exists and is not some image ending
	if !tld.IsValid(ending) {
		return false
	}

	if _, err := strconv.Atoi(ending); err == nil {
		return false
	}

	return true
}
