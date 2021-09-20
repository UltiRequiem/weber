package internal

import (
	"fmt"
	"log"
	"net/http"
)

func Fetch(url string, c chan bool) {
	resp, err := http.Get(url)

	if err != nil {
		return
	}

	resp.Body.Close()
	c <- true
}

func CallAllGoroutines(times int, channel chan bool, logFetch bool, url string) {
	for i := 0; i < times; i++ {
		if logFetch && <-channel {
			log.Println(fmt.Sprintf("Fetch %d to %s done successfully!", i+1, url))
		}
	}
}

func CicleFetch(times int, url string, channel chan bool) {
	for gophers := 0; gophers < times; gophers++ {
		go Fetch(url, channel)
	}
}

func CheckURL(url string, logCheck bool) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Error while fetching, check the url.")
	}

	if logCheck {
		log.Println(fmt.Sprintf("Status Code: %d", resp.StatusCode))
	}

	resp.Body.Close()
}
