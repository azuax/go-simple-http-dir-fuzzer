package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"

	"github.com/fatih/color"
)

func checkURL(baseURL string, urlChan chan string, WG *sync.WaitGroup) {
	defer WG.Done()

	path := <-urlChan
	re := regexp.MustCompile(`^/`)
	path = re.ReplaceAllString(path, "")
	fullURL := fmt.Sprintf("%s/%s", baseURL, path)

	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Printf("[Error] %s\n", fullURL)
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		color.Green(fmt.Sprintf("[%d] %s\n", resp.StatusCode, fullURL))
	} else {
		color.Red(fmt.Sprintf("[%d] %s\n", resp.StatusCode, fullURL))
	}

}

func main() {
	WG := new(sync.WaitGroup)
	params := ParseOpts(os.Args)
	urlChan := make(chan string, params.nThreads)
	fmt.Printf("We are going to run over %s with %d goroutines \n", params.URL, params.nThreads)

	f, err := os.Open(params.wl)
	if err != nil {
		log.Fatalln("Can't open wordlist file")
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		urlChan <- scanner.Text()
		WG.Add(1)
		go checkURL(params.URL, urlChan, WG)
	}

	WG.Wait()
}
