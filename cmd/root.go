package cmd

import (
	"fmt"
	"github.com/UltiRequiem/weber/internal"
	"log"
	"math"
	"time"
)

func Init() {
	channel := make(chan bool)

	url, logFetch, times, timeToSleep, maxChunkValue := getParams()

	if url == "" {
		log.Fatal("You have to pass an url!")
	}

	if logFetch {
		fmt.Println(fmt.Sprintf(`URL to Fetch: "%s", Times to fetch: %d, Log: %t.`, url, times, logFetch))
	}

	if internal.TestUrl(url, logFetch) {

		if times > maxChunkValue {
			timesToRepeatBucle := math.Round(float64(times / maxChunkValue))

			for i := 0.0; i < timesToRepeatBucle; i++ {
				for gophers := 0; gophers < 100; gophers++ {
					go internal.Fetch(url, channel)
				}

			}

		} else {
			for gophers := 0; gophers < times; gophers++ {
				go internal.Fetch(url, channel)
			}
		}

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
