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

	url, logFetch, times, timeToSleep, maxChunkCapacity := getParams()

	validateURL(url, logFetch)

	if logFetch {
		log.Println(fmt.Sprintf(`URL to Fetch: "%s", Log: %t, Times: %d, Max Chunk Capacity:%d.`, url, logFetch, times, maxChunkCapacity))
		defer log.Println("Process Finished.")
	}

	if times > maxChunkCapacity {
		timesToRepeatBucle := int(math.Round(float64(times / maxChunkCapacity)))

		channelOfChannels := make(chan string)
		for i := 0; i < timesToRepeatBucle; i++ {
			go func() {
				internal.CicleFetch(maxChunkCapacity, url, channel)
				channelOfChannels <- fmt.Sprintf("Chunk %d sended!", i+1)
			}()

			log.Println(<-channelOfChannels)
			time.Sleep(time.Second / time.Duration(timeToSleep))
		}

	} else {
		internal.CicleFetch(times, url, channel)
	}

	internal.CallAllGoroutines(times, channel, logFetch, url)

}
