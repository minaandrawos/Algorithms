/*
  String combinations in Go

  This is an exercise of recursion typically covered in programming training books:
  PROBLEM: Implement a function that prints all possible combinations of the
  characters in a string. These combinations range in length from one to the length
  of the string. Two combinations that differ only in ordering of their characters are
  the same combination. In other words, “12” and “31” are different combinations
  from the input string “123”, but “21” is the same as “12”.

*/
package main

import "fmt"

type Combinations struct {
	out []rune
	in  []rune
}

func startCombinations(input string) {
	c := &Combinations{
		in: []rune(input),
	}

	c.Combine(0)
}

func (c *Combinations) Combine(seed int) {
	inLen := len(c.in)
	for i := seed; i < inLen-1; i++ {
		c.out = append(c.out, c.in[i])
		fmt.Println(string(c.out))
		c.Combine(i + 1)
		c.out = c.out[:len(c.out)-1]
	}
	c.out = append(c.out, c.in[inLen-1])
	fmt.Println(string(c.out))
	c.out = c.out[:len(c.out)-1]
}
