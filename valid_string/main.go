/*
	Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

	An input string is valid if:

	Open brackets must be closed by the same type of brackets.
	Open brackets must be closed in the correct order.


	Example 1:
	Input: s = "()"
	Output: true

	Example 2:
	Input: s = "()[]{}"
	Output: true

	Example 3:
	Input: s = "(]"
	Output: false

	Example 4:
	Input: s = "([)]"
	Output: false

	Example 5:
	Input: s = "{[]}"
	Output: true
*/

package main

import (
	"fmt"
)

var maps = map[rune]rune{
	')': '(',
	'}': '{',
	']': '[',
}

func isValid(s string) bool {
	queue := []rune{}

	for _, c := range s {
		switch c {
		case '(', '{', '[':
			queue = append(queue, c)
		case ')', '}', ']':
			if queue[len(queue)-1] != maps[c] {
				return false
			}
			queue = queue[0 : len(queue)-1]
		}
	}

	return len(queue) == 0
}

func main() {
	type test struct {
		s    string
		want bool
	}

	tests := []test{
		{
			"()",
			true,
		},
		{
			"()[]{}",
			true,
		},
		{
			"([{}])",
			true,
		},
		{
			"(]",
			false,
		},
		{
			"([)]",
			false,
		},
		{
			"{[]}",
			true,
		},
	}

	for i, tt := range tests {
		got := isValid(tt.s)
		res := "PASSED"
		if got != tt.want {
			res = "FAILED"
		}
		fmt.Printf("\n%s case=%d got=%v want=%v", res, i, got, tt.want)
	}
}
