package cmd

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func fetch(url string, c chan bool) {
	_, err := http.Get(url)

	if err != nil {
		switch err.(type) {
		default:
			log.Fatal(err)
		}
	}

	c <- true
}

func Init() {
	channel := make(chan bool)

	url, logFetch, times := getParams()

	if logFetch {
		fmt.Println(fmt.Sprintf(`URL to Fetch: "%s", Times to fetch: %d, Log: %t.`, url, times, logFetch))
	}

	// Send Gophers to fetch the URL
	for i := 0; i < times; i++ {
		go fetch(url, channel)
	}

	// Call All the Gophers
	for i := 0; i < times; i++ {
		time.Sleep(time.Second / 2)

		if logFetch && <-channel {
			log.Println(fmt.Sprintf("%s fetched successfully!", url))
		}
	}

	if logFetch {
		log.Println("Finished")
	}
}
