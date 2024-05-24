package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	websites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.com",
		"https://amazon.com",
		"https://some-random-website-that-does-not-exists.eu",
	}

	for _, website := range websites {
		fmt.Printf("Checking: %v\n", website)
		resp, err := http.Get(website)
		if err != nil {
			log.Printf("Error with GET request to %v: %v", website, err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Printf("Unexpected status code for %v: %v\n", website, resp.StatusCode)
		} else {
			log.Println("All is good")
		}

	}
}
