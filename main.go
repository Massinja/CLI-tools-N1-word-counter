package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool) int {
	// scanner is used to read text from a reader
	scanner := bufio.NewScanner(r)

	if !countLines {
		// split scan to words (default: lines)
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	// add a flag to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	flag.Parse()

	fmt.Println(count(os.Stdin, *lines))

}
