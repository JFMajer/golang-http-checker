package main

import (
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
		log.Info().Msgf("Checking: %v", website)

		statusCode, err := statusCheck(website)
		if err != nil {
			log.Error().Err(err).Msgf("Error with GET request to %v", website)
			continue
		}

		if statusCode != http.StatusOK {
			log.Error().Int("status_code", statusCode).Msgf("Unexpected status code for %v", website)
		} else {
			log.Info().Msgf("All is good with %v", website)
		}

	}
}

func statusCheck(url string) (int, error) {
	log.Trace().Msgf("Making GET request to %v", url)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	log.Debug().Int("status_code", resp.StatusCode).Msgf("Received response for %v", url)
	return resp.StatusCode, nil
}
