package cmd

import (
	"github.com/UltiRequiem/weber/internal"
	"log"
)

func validateURL(url string, logFetch bool) {
	if url == "" {
		log.Fatal("You have to pass an URL!")
	}

	internal.TestUrl(url, logFetch)
}
