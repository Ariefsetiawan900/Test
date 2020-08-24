// Answer no. 3
// Used  find missing number function as an output.

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", answer)
	http.ListenAndServe(":8080", nil)
}

func findMissingNumber(arr []int, n int) int {
	sum := 0
	nNumSum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	nNumSum = ((n + 1) * (n + 2)) / 2
	return (nNumSum - sum)
}

func answer(w http.ResponseWriter, r *http.Request) {
	// You can change this value manually if you want to check the find missing number function.
	arr := []int{1, 2, 3, 4, 5, 6, 7, 9}

	size1 := len(arr)

	missingNo := findMissingNumber(arr, size1)
	fmt.Fprintf(w, "missing number in the array of size %d is %d ", size1+1, missingNo)

}
