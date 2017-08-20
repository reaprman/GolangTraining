package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	fmt.Println("------------------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("------------------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len: ", len(mySlice), "Capacity: ", cap(mySlice), "Value: ", mySlice[i])
	}
}

/*
Run Results:
------------------------
[]
0
5
------------------------
Len:  1 Capacity:  5 Value:  0
Len:  2 Capacity:  5 Value:  1
Len:  3 Capacity:  5 Value:  2
Len:  4 Capacity:  5 Value:  3
Len:  5 Capacity:  5 Value:  4
Len:  6 Capacity:  10 Value:  5
Len:  7 Capacity:  10 Value:  6
Len:  8 Capacity:  10 Value:  7
Len:  9 Capacity:  10 Value:  8
Len:  10 Capacity:  10 Value:  9
Len:  11 Capacity:  20 Value:  10
Len:  12 Capacity:  20 Value:  11
Len:  13 Capacity:  20 Value:  12
Len:  14 Capacity:  20 Value:  13
Len:  15 Capacity:  20 Value:  14
Len:  16 Capacity:  20 Value:  15
Len:  17 Capacity:  20 Value:  16
Len:  18 Capacity:  20 Value:  17
Len:  19 Capacity:  20 Value:  18
Len:  20 Capacity:  20 Value:  19
Len:  21 Capacity:  40 Value:  20
Len:  22 Capacity:  40 Value:  21
Len:  23 Capacity:  40 Value:  22
Len:  24 Capacity:  40 Value:  23
Len:  25 Capacity:  40 Value:  24
Len:  26 Capacity:  40 Value:  25
Len:  27 Capacity:  40 Value:  26
Len:  28 Capacity:  40 Value:  27
Len:  29 Capacity:  40 Value:  28
Len:  30 Capacity:  40 Value:  29
Len:  31 Capacity:  40 Value:  30
Len:  32 Capacity:  40 Value:  31
Len:  33 Capacity:  40 Value:  32
Len:  34 Capacity:  40 Value:  33
Len:  35 Capacity:  40 Value:  34
Len:  36 Capacity:  40 Value:  35
Len:  37 Capacity:  40 Value:  36
Len:  38 Capacity:  40 Value:  37
Len:  39 Capacity:  40 Value:  38
Len:  40 Capacity:  40 Value:  39
Len:  41 Capacity:  80 Value:  40
Len:  42 Capacity:  80 Value:  41
Len:  43 Capacity:  80 Value:  42
Len:  44 Capacity:  80 Value:  43
Len:  45 Capacity:  80 Value:  44
Len:  46 Capacity:  80 Value:  45
Len:  47 Capacity:  80 Value:  46
Len:  48 Capacity:  80 Value:  47
Len:  49 Capacity:  80 Value:  48
Len:  50 Capacity:  80 Value:  49
Len:  51 Capacity:  80 Value:  50
Len:  52 Capacity:  80 Value:  51
Len:  53 Capacity:  80 Value:  52
Len:  54 Capacity:  80 Value:  53
Len:  55 Capacity:  80 Value:  54
Len:  56 Capacity:  80 Value:  55
Len:  57 Capacity:  80 Value:  56
Len:  58 Capacity:  80 Value:  57
Len:  59 Capacity:  80 Value:  58
Len:  60 Capacity:  80 Value:  59
Len:  61 Capacity:  80 Value:  60
Len:  62 Capacity:  80 Value:  61
Len:  63 Capacity:  80 Value:  62
Len:  64 Capacity:  80 Value:  63
Len:  65 Capacity:  80 Value:  64
Len:  66 Capacity:  80 Value:  65
Len:  67 Capacity:  80 Value:  66
Len:  68 Capacity:  80 Value:  67
Len:  69 Capacity:  80 Value:  68
Len:  70 Capacity:  80 Value:  69
Len:  71 Capacity:  80 Value:  70
Len:  72 Capacity:  80 Value:  71
Len:  73 Capacity:  80 Value:  72
Len:  74 Capacity:  80 Value:  73
Len:  75 Capacity:  80 Value:  74
Len:  76 Capacity:  80 Value:  75
Len:  77 Capacity:  80 Value:  76
Len:  78 Capacity:  80 Value:  77
Len:  79 Capacity:  80 Value:  78
Len:  80 Capacity:  80 Value:  79
*/
