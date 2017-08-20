package main

import "fmt"

func main() {
	var x [256]int

	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 0; i < 256; i++ {
		x[i] = i
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
0 - int - 0
1 - int - 1
2 - int - 10
3 - int - 11
4 - int - 100
5 - int - 101
6 - int - 110
7 - int - 111
8 - int - 1000
9 - int - 1001
10 - int - 1010
11 - int - 1011
12 - int - 1100
13 - int - 1101
14 - int - 1110
15 - int - 1111
16 - int - 10000
17 - int - 10001
18 - int - 10010
19 - int - 10011
20 - int - 10100
21 - int - 10101
22 - int - 10110
23 - int - 10111
24 - int - 11000
25 - int - 11001
26 - int - 11010
27 - int - 11011
28 - int - 11100
29 - int - 11101
30 - int - 11110
31 - int - 11111
32 - int - 100000
33 - int - 100001
34 - int - 100010
35 - int - 100011
36 - int - 100100
37 - int - 100101
38 - int - 100110
39 - int - 100111
40 - int - 101000
41 - int - 101001
42 - int - 101010
43 - int - 101011
44 - int - 101100
45 - int - 101101
46 - int - 101110
47 - int - 101111
48 - int - 110000
49 - int - 110001
50 - int - 110010
51 - int - 110011
*/