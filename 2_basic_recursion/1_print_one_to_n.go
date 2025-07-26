package basic_recursion

import (
	"bytes"
	"fmt"
	"io"
	"os"
	
)

/*
Q. Print 1 to N using Recursion

Given an integer n, write a function to print all numbers from 1 to n (inclusive) using recursion.

You must not use any loops such as for, while, or do-while.
The function should print each number on a separate line, in increasing order from 1 to n.

Examples:
Input: n = 5

Output:
1  
2  
3  
4  
5

Input: n = 1

Output:
1

Constraints:
1 <= n <= 100
*/

func printOneToN(input int) {
	numberPrinter(1, input)
}
func numberPrinter(index int, input int) {
	if index > input {
		return
	}
	fmt.Println(index)
	numberPrinter(index+1, input)
}

var stdout io.Writer = os.Stdout

func TestPrintOneToN() {
	testCases := []struct {
		input    int
		expected string
	}{
		{5, "1\n2\n3\n4\n5"},
		{1, "1"},
		{3, "1\n2\n3"},
		{10, "1\n2\n3\n4\n5\n6\n7\n8\n9\n10"},
		{7, "1\n2\n3\n4\n5\n6\n7"},
		{2, "1\n2"},
		{4, "1\n2\n3\n4"},
		{6, "1\n2\n3\n4\n5\n6"},
		{8, "1\n2\n3\n4\n5\n6\n7\n8"},
		{9, "1\n2\n3\n4\n5\n6\n7\n8\n9"},
	}

	for _, tc := range testCases {
		// Capture output
		var buf bytes.Buffer
		oldStdout := stdout
		stdout = &buf
		defer func() { stdout = oldStdout }()
 fmt.Printf("Input :%d\n",tc.input)
		printOneToN(tc.input)

		// Compare output
		// result := strings.TrimSpace(buf.String())
		// if result == tc.expected {
		// 	println("PASS:", tc.input)
		// } else {
		// 	println("FAIL:", tc.input)
		// 	println("Expected:\n" + tc.expected)
		// 	println("Got:\n" + result)
		// }
	}
}
