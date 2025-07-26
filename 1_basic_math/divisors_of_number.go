package basic_math

func DivisorOfNumber(number int) []int {
	result := []int{1}
	for i := 2; i <= number; i++ {
		if number%i == 0 {
			result = append(result, i)
		}
	}
	return result
}

func TestDivisorOfNumber() {
	testCases := []struct {
		input    int
		expected []int
	}{
		{6, []int{1, 2, 3, 6}},
		{8, []int{1, 2, 4, 8}},
		{12, []int{1, 2, 3, 4, 6, 12}},
		{16, []int{1, 2, 4, 8, 16}},
		{7, []int{1, 7}}, // Prime number
		{1, []int{1}},    // Edge case
		{20, []int{1, 2, 4, 5, 10, 20}},
		{15, []int{1, 3, 5, 15}},
		{28, []int{1, 2, 4, 7, 14, 28}},
		{100, []int{1, 2, 4, 5, 10, 20, 25, 50, 100}},
	}

	for _, tc := range testCases {
		result := DivisorOfNumber(tc.input)
		equal := len(result) == len(tc.expected)
		if equal {
			for i := range result {
				if result[i] != tc.expected[i] {
					equal = false
					break
				}
			}
		}

		if equal {
			println("PASS:", tc.input)
		} else {
			print("FAIL: Input:", tc.input, " Expected: [")
			for i, v := range tc.expected {
				if i > 0 {
					print(" ")
				}
				print(v)
			}
			print("] Got: [")
			for i, v := range result {
				if i > 0 {
					print(" ")
				}
				print(v)
			}
			println("]")
		}
	}
}
