// ********RoostGPT********
/*
Test generated by RoostGPT for test go-test-examples using AI Type Open AI and AI Model gpt-4

ROOST_METHOD_HASH=Sum_d25bf13255
ROOST_METHOD_SIG_HASH=Sum_07f860fff8

Existing Test Information:
These test cases are already implemented and not included for test generation scenario:
File: go-test-examples/03-assertion/assert_test.go
Test Cases:
    [TestSum
    TestSumCustomAssertion
    TestSumTable]

Note: Only generate test cases based on the given scenarios,do not generate test cases other than these scenarios
Scenario 1: {Description:Test scenario where both input numbers are negative integers
Scenario 2: selected:true}
Scenario 3: {Description:Test scenario where one input number is positive and the other is negative
Scenario 4: {Description:Test scenario where both input numbers are zero
*/

// ********RoostGPT********
package main

import (
	"errors"
)

func Sum(x, y int) (int, error) {
	if x < 0 && y < 0 {
		return 0, errors.New("Both numbers are negative")
	}
	z := x + y
	return z, nil
}

func main() {
	Sum(5, 3)
}
