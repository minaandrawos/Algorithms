/*
	Selection sort implementation in Go
*/
package main

func findMinimumIndex(array []int, startIndex int) int {
	if startIndex > len(array)-1 {
		return -1
	}
	retVal := startIndex
	min := array[startIndex]
	for i := startIndex + 1; i <= len(array)-1; i++ {
		if min > array[i] {
			min = array[i]
			retVal = i
		}
	}
	return retVal
}

func selectionSort(array *[]int) {
	myArray := *array
	for i := 0; i <= len(myArray)-1; i++ {
		m := findMinimumIndex(myArray, i)
		temp := myArray[i]
		myArray[i] = myArray[m]
		myArray[m] = temp
	}
}
