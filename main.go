package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	websites := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.com",
		"https://some-random-website-that-does-not-exists.eu",
		"https://amazon.com",
	}

	for _, website := range websites {
		fmt.Printf("Checking: %v\n", website)
		resp, err := http.Get(website)
		if err != nil {
			log.Error().Err(err).Msgf("Error with GET request to %v", website)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Error().Int("status_code", resp.StatusCode).Msgf("Unexpected status code for %v", website)
		} else {
			log.Info().Msgf("All is good with %v", website)
		}

	}
}
