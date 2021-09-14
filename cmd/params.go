package cmd

import "flag"

func getParams() (string, bool, int, int64) {
	url := flag.String("url", "", "The URL to Fetch")
	times := flag.Int("times", 1, "How many times to fetch")
	timeToSleep := flag.Int64("timeToSleep", 2, "How many seconds to sleep.")
	logFetch := flag.Bool("log", false, "Log all the process.")

        // Shorter Versions
	u := flag.String("u", "", "The URL to Fetch")
	t := flag.Int("t", 1, "How many times to fetch")
	ts := flag.Int64("ts", 2, "How many seconds to sleep.")
	l := flag.Bool("l", false, "Log all the process.")

	flag.Parse()

	if *url == "" {
		*url = *u

	}

	if *times == 1 {
		*times = *t
	}

	if *timeToSleep == 2 {
		*timeToSleep = *ts
	}

	if *logFetch == false {
		*logFetch = *l
	}

	return *url, *logFetch, *times, *timeToSleep
}
