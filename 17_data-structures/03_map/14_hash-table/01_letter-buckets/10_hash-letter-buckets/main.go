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
	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// create slice to hold counts
	buckets := make([]int, 200)
	// Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:123])
	// fmt.Println("******************")
	// for i := 28; i <- 126; i++ {
	// fmt.Printf("%v - %c - %v \n", i , i, buckets[i])
	// }
}

// HashBucket - creats hash
func HashBucket(word string) int {
	return int(word[0])
}

/*
Run Result:
[1790 1260 910 522 325 658 473 812 2829 244 80 504 563 672 433 816 291 259 1
671 1771 87 119 1080 5 205 15 21 0 0 0 2 0 21787 9692 7385 5219 3614 7576 30
16 12620 11607 582 852 5321 7765 4060 13798 5256 395 3592 16260 33335 2591 1
452 13314 1 2272 19]
*/
