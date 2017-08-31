package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book adventures of moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}
	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// set the split function for the scnning operation.
	scanner.Split(bufio.ScanWords)
	// create slice to hold counts
	buckets := make([]int, 12)
	// loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets)
}

// HashBucket - Hash function
func HashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}

/*
Run Result:
[17392 26570 12820 17065 15797 12809 11729 21409 13999 32182 13037 19303]
*/
