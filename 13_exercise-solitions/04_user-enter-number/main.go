package main

import "fmt"

func main() {
	var nHigh, nLow int

	fmt.Print("Please enter the large number: ")
	fmt.Scan(&nHigh)
	fmt.Print("Enter the lower number: ")
	fmt.Scan(&nLow)
	if nLow > nHigh {
		fmt.Println("The number entered is invalid: ")
		fmt.Scan(&nLow)
	}

	fmt.Println("The remainder is: ", nHigh%nLow)
	fmt.Println("The quotient is: ", nHigh/nLow)
}

/*
4. Create a program that print to the terminal asking for a
user to enter a small number and a larger number. Print the
remainder of the larger number divided by the smaller number

Run Results:
Please enter the large number: 58
Enter the lower number: 8
The remainder is:  2
The quotient is:  7
*/
