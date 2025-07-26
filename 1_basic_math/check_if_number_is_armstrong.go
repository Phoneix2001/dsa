package basic_math

import (
	"math"
	"strconv"
)

func IsArmstrong(a int) bool {
	var total int
	var length int
	var comparing = a
	if a < 0 {
		length = -a
	} else {
		length = a
	}
	length = len(strconv.Itoa(length))
	for a >= 1 {
		
		digit := a % 10
		a = a / 10
		total += int(math.Pow(float64(digit), float64(length)))
	}
	
	if total == comparing {
		return true
	} else {
		return false
	}
}

func TestIsArmstrong() {
	testCases := []struct {
		input    int
		expected bool
	}{
		{153, true},
		{370, true},
		{371, true},
		{407, true},
		{0, true},
		{1, true},
		{10, false},
		{9474, true},
		{9475, false},
		{947, false},
	}

	for _, tc := range testCases {
		result := IsArmstrong(tc.input)
		if result == tc.expected {
			println("PASS:", tc.input)
		} else {
			println("FAIL:", tc.input, "Expected:", tc.expected, "Got:", result)
		}
	}
}
