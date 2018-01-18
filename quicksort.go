package main

func quickSort(input []int, left, right int) {
	index := partition(input, left, right)
	if index-1 > left {
		quickSort(input, left, index-1)
	}
	if index < right {
		quickSort(input, index, right)
	}
}

func partition(input []int, left, right int) int {
	pivot := input[(left+right)/2]
	for left <= right {
		for pivot > input[left] {
			left++
		}

		for pivot < input[right] {
			right--
		}

		if left <= right {
			temp := input[left]
			input[left] = input[right]
			input[right] = temp
			left++
			right--
		}
	}
	return left
}
