package cmd

import "flag"

func getParams() (string, bool, int, int64) {
	url := flag.String("url", "", "The URL to Fetch")
	times := flag.Int("t", 1, "How many times to fetch")

	timeToSleep := flag.Int64("ts", 2, "How many times to fetch")

	logFetch := flag.Bool("l", false, "Log when a site is fetched")

	flag.Parse()

	return *url, *logFetch, *times, *timeToSleep
}
