/*
  This algorithms can retrieve the minimum k elements of an array, it runs at O(KN) time.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	result := getTopK([]int{8, 9, 10, 11, 32, 44, 2, 1, 4}, 6)
	fmt.Println(result)
}

func getTopK(array []int, k int) []int {
	result := make([]int, k)
	for i := range result {
		result[i] = math.MaxInt64
	}

	for _, v := range array {
		current := v
		for i := 0; i < k; i++ {
			if result[i] == math.MaxInt64 {
				result[i] = current
				break
			}

			if result[i] > current {
				temp := result[i]
				result[i] = current
				current = temp
			}
		}
	}
	return result
}
