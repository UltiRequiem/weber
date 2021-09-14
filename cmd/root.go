package cmd

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func fetch(url string, c chan bool) {
	resp, _ := http.Get(url)
	resp.Body.Close()
	c <- true
}

func testUrl(url string, logCheck bool) bool {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Error while fetching, check the url.")
	}

	if logCheck {
		log.Println(fmt.Sprintf("Status Code: %d", resp.StatusCode))
	}

	return true

}

func Init() {
	channel := make(chan bool)

	url, logFetch, times, timeToSleep := getParams()

	if logFetch {
		fmt.Println(fmt.Sprintf(`URL to Fetch: "%s", Times to fetch: %d, Log: %t.`, url, times, logFetch))
	}

	if testUrl(url, logFetch) {

		// Send Gophers to fetch the URL
		for i := 0; i < times; i++ {
			go fetch(url, channel)
		}

		// Call All the Gophers
		for i := 0; i < times; i++ {
			time.Sleep(time.Second / time.Duration(timeToSleep))

			if logFetch && <-channel {
				log.Println(fmt.Sprintf("%s fetched successfully!", url))
			}
		}
	}

	if logFetch {
		log.Println("Finished")
	}
}
