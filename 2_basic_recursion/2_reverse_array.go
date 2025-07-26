package basic_recursion


func reverseArray(input []int) {
reversingArray(input,len(input)-1,0)
}

func reversingArray(input []int,len int,index int) {
if index > len {
	return
}
input[index] ,input[len] = input[len],input[index]
reversingArray(input,len-1,index+1)
}


func TestReverseArray() {
	testCases := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Basic 5 elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "Array with duplicates",
			input:    []int{1, 2, 1, 1, 5, 1},
			expected: []int{1, 5, 1, 1, 2, 1},
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Two elements",
			input:    []int{1, 2},
			expected: []int{2, 1},
		},
		{
			name:     "All same elements",
			input:    []int{1, 1, 1, 1},
			expected: []int{1, 1, 1, 1},
		},
		{
			name:     "Negative numbers",
			input:    []int{-1, -2, -3, -4},
			expected: []int{-4, -3, -2, -1},
		},
		{
			name:     "Mixed numbers",
			input:    []int{-1, 0, 1, -2, 2},
			expected: []int{2, -2, 1, 0, -1},
		},
		{
			name:     "Large numbers",
			input:    []int{1000, 2000, 3000},
			expected: []int{3000, 2000, 1000},
		},
		{
			name:     "Zero array",
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
	}

	for _, tc := range testCases {
		// Create a copy of input array to avoid modifying the original
		inputCopy := make([]int, len(tc.input))
		copy(inputCopy, tc.input)

		reverseArray(inputCopy)

		// Check if arrays are equal
		isEqual := len(inputCopy) == len(tc.expected)
		if isEqual {
			for i := range inputCopy {
				if inputCopy[i] != tc.expected[i] {
					isEqual = false
					break
				}
			}
		}

		if isEqual {
			println("PASS:", tc.name)
		} else {
			print("FAIL:", tc.name, "\nInput: [")
			printArray(tc.input)
			print("]\nExpected: [")
			printArray(tc.expected)
			print("]\nGot: [")
			printArray(inputCopy)
			println("]")
		}
	}
}

// Helper function to print arrays
func printArray(arr []int) {
	for i, v := range arr {
		if i > 0 {
			print(" ")
		}
		print(v)
	}
}
