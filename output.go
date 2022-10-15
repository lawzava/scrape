package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	outputPlain = "plain"
	outputCSV   = "csv"
	outputJSON  = "json"
)

//nolint:forbidigo // allow println here for non-intrusive response
func handleOutput(emails []string) {
	switch output {
	case outputCSV:
		for _, email := range emails {
			if outputWithURL {
				fmt.Print(url + ",")
			}

			fmt.Println(email)
		}
	case outputJSON:
		fmt.Println(prepareJSONOutput(emails))
	default:
		for _, email := range emails {
			if outputWithURL {
				fmt.Print(url + " ")
			}

			fmt.Println(email)
		}
	}
}

func prepareJSONOutput(emails []string) string {
	if outputWithURL {
		type outputFormat struct {
			URL   string `json:"url"`
			Email string `json:"email"`
		}

		out := make([]outputFormat, len(emails))
		for i, email := range emails {
			out[i].URL = url
			out[i].Email = email
		}

		b, err := json.Marshal(out)
		if err != nil {
			log.Fatal(fmt.Errorf("failed to marshal json with url response: %w", err))
		}

		return string(b)
	}

	type outputFormat struct {
		Email string `json:"email"`
	}

	out := make([]outputFormat, len(emails))
	for i, email := range emails {
		out[i].Email = email
	}

	b, err := json.Marshal(out)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to marshal json response: %w", err))
	}

	return string(b)
}
