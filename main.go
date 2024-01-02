package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader) int {
	// scanner is used to read text from a reader
	scanner := bufio.NewScanner(r)

	// split scan to words (default: lines)
	scanner.Split(bufio.ScanWords)

	wc := 0

	for scanner.Scan() {
		wc++
	}

	return wc
}

func main() {
	fmt.Println(count(os.Stdin))

}
