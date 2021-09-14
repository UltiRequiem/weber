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

func TestUrl(url string, logCheck bool) bool {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal("Error while fetching, check the url.")
	}

	if logCheck {
		log.Println(fmt.Sprintf("Status Code: %d", resp.StatusCode))
	}

	resp.Body.Close()

	return true

}
