package main

import (
	"flag"
	"fmt"
	"os"
)

type Params struct {
	URL      string
	wl       string
	nThreads int
}

func ParseOpts(args []string) (PARAMS Params) {
	if len(args) < 3 {
		fmt.Printf("Usage %s -u <URL> -w <WORDLIST> [-t <threads>]\n", args[0])
		os.Exit(1)
	}
	flag.StringVar(&PARAMS.URL, "u", "", "URL to test. Example: https://www.github.com")
	flag.StringVar(&PARAMS.wl, "w", "", "Wordlist to iterate. Example: /tmp/wordlist.txt")
	flag.IntVar(&PARAMS.nThreads, "t", 20, "Max number of concurrent requests")
	flag.Parse()

	return
}
