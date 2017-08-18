package main

import "fmt"

func main() {
	var x [256]byte
	fmt.Println(len(x))
	fmt.Println(x[0])
	for i := 0; i < 256; i++ {
		x[i] = byte(i)
	}
	for i, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}
}

/*
Run Results:
256
0
0 - uint8 - 0
1 - uint8 - 1
2 - uint8 - 10
3 - uint8 - 11
4 - uint8 - 100
5 - uint8 - 101
6 - uint8 - 110
7 - uint8 - 111
8 - uint8 - 1000
9 - uint8 - 1001
10 - uint8 - 1010
11 - uint8 - 1011
12 - uint8 - 1100
13 - uint8 - 1101
14 - uint8 - 1110
15 - uint8 - 1111
16 - uint8 - 10000
17 - uint8 - 10001
18 - uint8 - 10010
19 - uint8 - 10011
20 - uint8 - 10100
21 - uint8 - 10101
22 - uint8 - 10110
23 - uint8 - 10111
24 - uint8 - 11000
25 - uint8 - 11001
26 - uint8 - 11010
27 - uint8 - 11011
28 - uint8 - 11100
29 - uint8 - 11101
30 - uint8 - 11110
31 - uint8 - 11111
32 - uint8 - 100000
33 - uint8 - 100001
34 - uint8 - 100010
35 - uint8 - 100011
36 - uint8 - 100100
37 - uint8 - 100101
38 - uint8 - 100110
39 - uint8 - 100111
40 - uint8 - 101000
41 - uint8 - 101001
42 - uint8 - 101010
43 - uint8 - 101011
44 - uint8 - 101100
45 - uint8 - 101101
46 - uint8 - 101110
47 - uint8 - 101111
48 - uint8 - 110000
49 - uint8 - 110001
50 - uint8 - 110010
51 - uint8 - 110011
*/
