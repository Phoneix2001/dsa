package basic_math

import "math"

/*
Q. Divisors of a Number


You are given an integer n. You need to find all the divisors of n. Return all the divisors of n as an array or list in a sorted order.

A number which completely divides another number is called it's divisor.

Examples:
   Input: n = 6

   Output = [1, 2, 3, 6]

   Explanation: The divisors of 6 are 1, 2, 3, 6.

   Input: n = 8

   Output: [1, 2, 4, 8]

   Explanation: The divisors of 8 are 1, 2, 4, 8.
*/



func isPrime(input int) bool {
	if input < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(input))); i++ {
		if input != i && input%i == 0 {
			return false
		}
	}
	return true
}

func TestIsPrime() {
	testCases := []struct {
		input    int
		expected bool
	}{
		// First 25 prime numbers
		{2, true}, {3, true}, {5, true}, {7, true}, {11, true},
		{13, true}, {17, true}, {19, true}, {23, true}, {29, true},
		{31, true}, {37, true}, {41, true}, {43, true}, {47, true},
		{53, true}, {59, true}, {61, true}, {67, true}, {71, true},
		{73, true}, {79, true}, {83, true}, {89, true}, {97, true},

		// Common non-prime numbers
		{1, false}, {4, false}, {6, false}, {8, false}, {9, false},
		{10, false}, {12, false}, {14, false}, {15, false}, {16, false},
		{18, false}, {20, false}, {21, false}, {22, false}, {24, false},
		{25, false}, {26, false}, {27, false}, {28, false}, {30, false},

		// Larger prime numbers
		{101, true}, {103, true}, {107, true}, {109, true}, {113, true},
		{127, true}, {131, true}, {137, true}, {139, true}, {149, true},
		{151, true}, {157, true}, {163, true}, {167, true}, {173, true},

		// Larger non-prime numbers
		{100, false}, {111, false}, {119, false}, {121, false}, {129, false},
		{133, false}, {141, false}, {143, false}, {144, false}, {145, false},
		{146, false}, {147, false}, {148, false}, {150, false}, {152, false},

		// Edge cases and negative numbers
		{0, false}, {-1, false}, {-2, false}, {-3, false}, {-4, false},
		{-5, false}, {-6, false}, {-7, false}, {-8, false}, {-9, false},

		// Perfect squares and cubes (non-prime)
		{4, false}, {9, false}, {16, false}, {25, false}, {36, false},
		{49, false}, {64, false}, {81, false}, {100, false}, {121, false},

		// Numbers just before and after prime numbers
		{6, false}, {8, false}, {10, false}, {12, false}, {14, false},
		{16, false}, {18, false}, {20, false}, {22, false}, {24, false},
	}

	for _, tc := range testCases {
		result := isPrime(tc.input)
		if result == tc.expected {
			println("PASS:", tc.input)
		} else {
			println("FAIL:", tc.input, "Expected:", tc.expected, "Got:", result)
		}
	}
}
