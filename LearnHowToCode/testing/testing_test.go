package main

import "testing"

// Tests must be in file that ends with "_test.go"
// put the file in the same package as the one being tested
// be in a func with an signature "func TestXxx(*testing.T)
// run test: go test

func TestMySum(t *testing.T) {

	type test struct {
		data []int
		answer int
	}


	x :=  mySum(2,3)
	if x != 5 {
		t.Error("Expected", 5, "Got", x)
	}


	// table test
	tests :=[]test{
		test{[]int{21, 21}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{1, 9}, 10},
		test{[]int{-21, 12, 10}, 1},
		test{[]int{-1, 1, 0}, 0},
	}

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}

}