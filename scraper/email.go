package scraper

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/lawzava/go-tld"
)

type emails struct {
	emails []string
	m      sync.Mutex
}

func (s *emails) add(email string) {
	if !isValidEmail(email) {
		return
	}

	// check for already existing emails
	for _, existingEmail := range s.emails {
		if existingEmail == email {
			return
		}
	}

	s.m.Lock()
	s.emails = append(s.emails, email)
	s.m.Unlock()
}

// Initialize once.
var reg = regexp.MustCompile(`([a-zA-Z0-9._-]+@([a-zA-Z0-9_-]+\.)+[a-zA-Z0-9_-]+)`)

// Parse any *@*.* string and append to the slice.
func (s *emails) parseEmails(body []byte) {
	res := reg.FindAll(body, -1)

	for _, r := range res {
		s.add(string(r))
	}
}

func (s *emails) parseCloudflareEmail(cloudflareEncodedEmail string) {
	decodedEmail := decodeCloudflareEmail(cloudflareEncodedEmail)
	email := reg.FindString(decodedEmail)

	s.add(email)
}

func decodeCloudflareEmail(email string) string {
	var e bytes.Buffer

	r, _ := strconv.ParseInt(email[0:2], 16, 0)

	for n := 4; n < len(email)+2; n += 2 {
		i, _ := strconv.ParseInt(email[n-2:n], 16, 0)
		c := i ^ r

		e.WriteRune(rune(c))
	}

	return e.String()
}

// Check if email looks valid.
func isValidEmail(email string) bool {
	if email == "" {
		return false
	}

	split := strings.Split(email, ".")

	// nolint:gomnd // allow magic number here
	if len(split) < 2 {
		return false
	}

	ending := split[len(split)-1]

	// nolint:gomnd // allow magic number here
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
