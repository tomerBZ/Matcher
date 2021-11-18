package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sync"
	"time"

	"BigID/aggregator"
	"BigID/match"
)

const (
	chunksSize   = 90 * 1024
	fileLocation = "assets/big.txt"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()
	ag := aggregator.NewAggregator()
	matcher := match.NewMatcher(ag)
	f, err := os.Open(fileLocation)
	if err != nil {
		log.Fatal("not able to read the file", err)
		return
	}
	defer f.Close()
	r := bufio.NewReader(f)
	processFile(r, &wg, matcher)
	wg.Wait()
	log.Printf("%+v\n", ag.GetAggregation())
	log.Println("finished processing the file at: ", time.Since(start))
}

func processFile(
	r *bufio.Reader,
	wg *sync.WaitGroup,
	matcher match.Matcher,
) {
	for {
		// divide chunks to 1006 lines
		chunk := make([]byte, chunksSize)
		n, err := r.Read(chunk)
		chunk = chunk[:n]
		if n == 0 {
			break
		}
		// byte each line and append to chunk strings
		newline, err := r.ReadString('\n')
		if err != io.EOF {
			chunk = append(chunk, newline...)
		}

		wg.Add(1)
		go func() {
			go matcher.Find(chunk)
			wg.Done()
		}()
	}
}
