package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book adventures of sherlock holmes
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}
	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// creates slice of slice of string to hold slices of words
	buckets := make([][]string, 12)
	// create slice to hold words
	for i := 0; i < 12; i++ {
		buckets = append(buckets, []string{})
	}
	// loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}
	// print len of each bucket
	for i := 0; i < 12; i++ {
		fmt.Println(i, " - ", len(buckets[i]))
	}
	// print the words in one of the buckets
	//fmt.Println(buckets[6])
}

// HashBucket - hash function
func HashBucket(word string, bucket int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % bucket
}

/*
 1st Run Result:
  501(c)(3) identification 501(c)(3) numerous 1500 (801) date can web gbnewby
@pglaf.org Information upon can accessible outdated $5,000) important commit
ted complying regulating effort, requirements. accepting approach make conce
rning our check other http://pglaf.org/donate General Information originator
 only loose unless notice start our search make our new our new]

 2nd Run:
0  -  8517
1  -  15702
2  -  6818
3  -  8229
4  -  8172
5  -  6444
6  -  5833
7  -  11097
8  -  6687
9  -  14526
10  -  5939
11  -  9569

*/
