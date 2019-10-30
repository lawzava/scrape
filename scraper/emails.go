package scraper

import (
	"regexp"
	"strconv"
	"strings"
)

// Parse any *@*.* string and append to the slice
func parseEmails(body []byte, scrapedEmails *[]string) {
	reg := regexp.MustCompile(`([a-zA-Z0-9._-]+@([a-zA-Z0-9_-]+\.)+[a-zA-Z0-9_-]+)`)
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

// Ignore these endings
var suffixFilter = []string{
	"png",
	"jpg",
	"gif",
	"svg",
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

	for _, sf := range suffixFilter {
		if ending == sf {
			return false
		}
	}

	if _, err := strconv.Atoi(ending); err == nil {
		return false
	}

	return true
}
