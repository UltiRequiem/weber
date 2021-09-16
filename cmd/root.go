package cmd

import (
	"fmt"
	"github.com/UltiRequiem/weber/internal"
	"log"
	"math"
	"time"
)

func Init() {
	url, logFetch, times, timeToSleep, maxChunkCapacity := getParams()

	// If something fail, this will kill the program
	validateURL(url, logFetch)

	if logFetch {
		log.Println(fmt.Sprintf(`URL to Fetch: "%s", Log: %t, Times: %d, Max Chunk Capacity:%d.`, url, logFetch, times, maxChunkCapacity))
		defer log.Println("Process finished successfully.")
	}

	channel := make(chan bool)

	if times > maxChunkCapacity {
		timesToRepeatBucle := int(math.Round(float64(times / maxChunkCapacity)))

		c := make(chan string)
		for i := 0; i < timesToRepeatBucle; i++ {
			go func() {
				internal.CicleFetch(maxChunkCapacity, url, channel)
				c <- fmt.Sprintf("Chunk %d sended!", i+1)
			}()

			log.Println(<-c)
			time.Sleep(time.Second / time.Duration(timeToSleep))
		}

	} else {
		internal.CicleFetch(times, url, channel)
	}

	internal.CallAllGoroutines(times, channel, logFetch, url)
}
