/*
 A collection of algorithms implemented in Go
*/
package main

import "fmt"

func main() {
	//BinarySearch
	array := []int{1, 2, 3, 4, 5, 6, 6, 10, 15, 15, 17, 18, 19, 30}
	find := binarySearch(array, 10, 0, len(array)-1)
	fmt.Printf("Found the element by recursive binary search at index %d \n", find)
	find = iterBinarySearch(array, 10, 0, len(array)-1)
	fmt.Printf("Found the element by iterative binary search at index %d \n", find)

	//String permutation
	input := "hat"
	startPermutation(input)

	//String combinations
	input = "wxyz"
	startCombinations(input)

	//Selection sort
	array = []int{7, 1, 9, 0, 2, 14, 10, 11, 4, 13}
	fmt.Printf("Array before sorting %v \n", array)
	selectionSort(&array)
	fmt.Printf("Selection sorted array %v \n", array)

	//Insertion sort
	array = []int{7, 1, 9, 0, 2, 14, 10, 11, 4, 13}
	fmt.Printf("Array before sorting %v \n", array)
	InsertionSort(&array)
	fmt.Printf("Insertion sorted array %v \n", array)
}
