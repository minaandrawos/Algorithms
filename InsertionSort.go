/*
	Insertion sort implementation in Go
*/
package main

func InsertionSort(myArray []int) {
	for i := 1; i < len(myArray); i++ {
		val := myArray[i]
		j := i - 1
		for j >= 0 && val < myArray[j] {
			myArray[j+1] = myArray[j]
			j--
		}
		myArray[j+1] = val
	}
}
