package main

import (
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type result struct {
	url        string
	statusCode int
	err        error
}

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

	c := make(chan result)

	for _, website := range websites {
		go statusCheck(website, c)
	}

	for range websites {
		res := <-c
		if res.err != nil {
			log.Error().Err(res.err).Msgf("Error with GET request to %v", res.url)
		} else if res.statusCode != http.StatusOK {
			log.Error().Int("status_code", res.statusCode).Msgf("Unexpected status code for %v", res.url)
		} else {
			log.Info().Msgf("All is good with %v", res.url)
		}
	}

}

func statusCheck(url string, c chan result) {
	log.Trace().Msgf("Making GET request to %v", url)
	resp, err := http.Get(url)
	if err != nil {
		c <- result{
			statusCode: 0,
			url:        url,
			err:        err,
		}
		return
	}
	defer resp.Body.Close()

	log.Debug().Int("status_code", resp.StatusCode).Msgf("Received response for %v", url)
	c <- result{
		statusCode: resp.StatusCode,
		url:        url,
		err:        nil,
	}
}
