package main

import "fmt"

func twoSum(nums []int, target int) []int {
	check := make([]int, 2)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				check[0] = i
				check[1] = j
			}
		}
	}
	return check
}

func main() {
	target := 9
	//nums := []int{2,7,11,15}
	fmt.Println("Please find the sum of two numbers that equal ", target)
	fmt.Println("From the followin list: ", []int{2, 7, 11, 15})
	fmt.Println("Answer: ", twoSum([]int{2, 7, 11, 15}, target))

}

/*
Given an array of integers, return indices of
the two numbers such that they add up to a specific target

Run Result:
Please find the sum of two numbers that equal  9
From the followin list:  [2 7 11 15]
Answer:  [0 1]
*/
