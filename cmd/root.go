package cmd

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

func fetch(url string, c chan bool) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
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

	if url == "" {
		log.Fatal("You have to pass an url!")
	}

	if logFetch {
		fmt.Println(fmt.Sprintf(`URL to Fetch: "%s", Times to fetch: %d, Log: %t.`, url, times, logFetch))
	}

	if testUrl(url, logFetch) {

		if times > 100 {
			timesToRepeatBucle := math.Round(float64(times / 100))

			for i := 0.0; i < timesToRepeatBucle; i++ {
				for gophers := 0; gophers < 100; gophers++ {
					go fetch(url, channel)
				}

			}

		} else {
			// Send Gophers to fetch the URL
			for i := 0; i < times; i++ {
				go fetch(url, channel)
			}
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
