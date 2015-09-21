/*
	String permutations in Go
	This is an exercise in recursion encountered in programming training books:

	PROBLEM: Implement a routine that prints all possible orderings of the characters
	in a string. In other words, print all permutations that use all the characters from
	the original string. For example, given the string “hat”, your function should print
	the strings “tha”, “aht”, “tah”, “ath”, “hta”, and “hat”. Treat each character in the
	input string as a distinct character, even if it is repeated. Given the string “aaa”,
	your routine should print “aaa” six times. You may print the permutations in any
	order you choose.
*/
package main

import "fmt"

func StringPermutations(input string, firstIndex int, secondIndex int, result *[]string) {

	if firstIndex == len(input) {
		return
	}

	fmt.Printf("first index %d, second index %d \n", firstIndex, secondIndex)
	input = goCharSwap(input, firstIndex, secondIndex)
	*result = append(*result, input)

	if secondIndex+1 != firstIndex {
		secondIndex++
	} else {
		secondIndex += 2
	}

	if secondIndex >= len(input) {
		secondIndex = 0
		firstIndex++
	}

	StringPermutations(input, firstIndex, secondIndex, result)
}

func goCharSwap(s string, firstIndex int, secondIndex int) string {
	inputBytes := []rune(s)
	temp := inputBytes[firstIndex]
	inputBytes[firstIndex] = inputBytes[secondIndex]
	inputBytes[secondIndex] = temp
	return string(inputBytes)
}

func startPermutation(input string) {
	p := &Permutation{
		in:   []rune(input),
		used: make([]bool, len(input)),
	}
	p.permute()
}

type Permutation struct {
	used []bool
	out  []rune
	in   []rune
}

func (p *Permutation) permute() {
	if len(p.out) == len(p.in) {
		fmt.Println(string(p.out))
		return
	}

	for i := 0; i < len(p.in); i++ {
		if p.used[i] == true {
			continue
		}
		p.out = append(p.out, p.in[i])
		p.used[i] = true
		p.permute()
		p.used[i] = false
		p.out = p.out[:len(p.out)-1]
	}
}
