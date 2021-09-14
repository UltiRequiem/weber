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

			channelOfChannels := make(chan string)
			for i := 1.0; i < timesToRepeatBucle+1; i++ {

				go func() {
					for gophers := 0; gophers < maxChunkValue; gophers++ {
						go internal.Fetch(url, channel)
					}

					channelOfChannels <- fmt.Sprintf("Chunk %d sended!", int(i))
				}()

				log.Println(<-channelOfChannels)
				time.Sleep(time.Second / time.Duration(timeToSleep))
			}

		} else {
			for gophers := 0; gophers < times; gophers++ {
				go internal.Fetch(url, channel)
			}
		}

		for i := 1; i < times+1; i++ {

			if logFetch && <-channel {
				log.Println(fmt.Sprintf("Fetch %d to %s done successfully!", i, url))
			}
		}
	}

	if logFetch {
		log.Println("Finished")
	}
}
