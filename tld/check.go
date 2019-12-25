package tld

func IsValid(request string) bool {
	for _, validTLD := range availableTLDs() {
		if validTLD == request {
			return true
		}
	}

	return false
}
