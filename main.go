package main

import (
	"log"
	"os"
)

func main() {
	err := rootCmd.Execute()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
