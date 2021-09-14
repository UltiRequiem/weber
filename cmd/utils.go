package cmd

import "flag"

func getParams() (string, bool, int) {
	url := flag.String("url", "https://google.com", "The URL to Fetch")
	times := flag.Int("t", 1, "How many times to fetch")

	logFetch := flag.Bool("l", false, "Log when a site is fetched")

	flag.Parse()

	return *url, *logFetch, *times

}
