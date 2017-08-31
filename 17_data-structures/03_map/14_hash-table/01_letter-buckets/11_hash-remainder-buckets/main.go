package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book Moby Dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// Create slice to hold counts
	buckets := make([]int, 12)
	// loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets)
}

//HashBucket - hash function
func HashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}

/*
Run Result:
[7909 34850 14167 22516 11229 6585 13167 20653 47306 14912 4503 16315]
*/
