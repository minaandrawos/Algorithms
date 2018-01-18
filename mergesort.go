package main

func main() {

}

func mergesort(array, helper []int, start, end int) {
	if start >= end {
		return
	}
	mid := (start + end) / 2
	merge(array, helper, start, mid)
	merge(array, helper, mid, end)
}

func merge(array, helper []int, start, end int) {
	for i, v := range array {
		helper[i] = v
	}

	middle := (start + end) / 2
	current := start
	right := middle + 1
	for start <= middle && right < end {
		if helper[start] <= helper[right] {
			array[current] = helper[start]
			start++
		} else if helper[right] > helper[start] {
			array[current] = helper[right]
			right++
		}
		current++
	}

	remaining := middle - start
	for i := 0; i < remaining; i++ {
		array[i+current] = helper[start+i]
	}
}
