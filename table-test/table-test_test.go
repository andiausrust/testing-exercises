package main

import "testing"

func TestMySumWithTable(t *testing.T) {

	type test struct {
		data []int
		answer int
	}

	tests := []test{
		test{
			data: []int{21, 21},
			answer: 42},
		test{
			data: []int{10, 21},
			answer: 31},
		test{
			data: []int{11, 21},
			answer: 32},
		test{
			data: []int{21, 22},
			answer: 43},
	}

	for _,y := range tests {
		x := mySumTable(y.data...)
		if x != y.answer {
			t.Error("Expected", y.answer, "got", x)
		}
	}
}
