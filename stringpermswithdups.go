/*
  Find all the permutations of a string that has duplicate characters, ensure that no permutation is repeated twice
*/
package main

import "fmt"

func main() {
	freq := make(map[byte]int)
	input := "122"
	for _, v := range input {
		freq[byte(v)]++
	}
	permsWithDuplicates(input, []byte{}, freq)
}

func permsWithDuplicates(input string, subset []byte, freqtable map[byte]int) {
	if len(subset) >= len(input) {
		fmt.Println(string(subset))
		return
	}

	for k := range freqtable {
		count := freqtable[k]
		if count > 0 {
			freqtable[k] = count - 1
			permsWithDuplicates(input, append(subset, k), freqtable)
			freqtable[k] = count
		}
	}
}
