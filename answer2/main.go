// Answer no. 2
// about find missing number function
package main

import (
	"fmt"
	"sort"
)

func findMissingNumber(arr []int, n int) int {
	sum := 0
	nNumSum := 0
	for i := 0; i < n; i++ {
		sum += arr[i] // sum variable will contain the sum fo all integers of an array.
	}
	nNumSum = ((n + 1) * (n + 2)) / 2 // we use formula for obtaining sum of first n numbers.
	return (nNumSum - sum)            // return the missing number.
}

func result(arr []int) {
	sort.Ints(arr)    //sort accending
	size1 := len(arr) //count index array
	//looping array
	for _, item := range arr {
		fmt.Print(item)
	}
	missingNo := findMissingNumber(arr, size1)
	fmt.Printf("  missing number in the array of size %d is %d ", size1+1, missingNo) // Printing the missing number

}
func main() {

	result([]int{1, 2, 3, 4, 5, 6, 7, 9, 10})

}
