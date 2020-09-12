package main

import "testing"

func Test_min64(t *testing.T) {
	testcases := [][]int64{
		{10, 20, 10},
		{20, 10, 10},
		{20, 30, 20},
		{0, 1, 0},
	}
	for _, testcase := range testcases {
		x := testcase[0]
		y := testcase[1]
		expected := testcase[2]
		got := min64(x, y)
		if got != expected {
			t.Errorf("got: %v\nexpected: %v", got, expected)
		}
	}
}

func Test_max64(t *testing.T) {
	testcases := [][]int64{
		{10, 20, 20},
		{20, 10, 20},
		{20, 30, 30},
		{0, 1, 1},
	}
	for _, testcase := range testcases {
		x := testcase[0]
		y := testcase[1]
		expected := testcase[2]
		got := max64(x, y)
		if got != expected {
			t.Errorf("got: %v\nexpected: %v", got, expected)
		}
	}
}

func Test_min64_2(t *testing.T) {
	testcases := [][]int64{
		{10, 20, 30, 10, 20},
		{10, 30, 20, 10, 20},
		{20, 10, 30, 10, 20},
		{20, 30, 10, 10, 20},
		{30, 10, 20, 10, 20},
		{30, 20, 10, 10, 20},
		{10, 10, 20, 10, 10},
		{10, 20, 10, 10, 10},
		{20, 10, 10, 10, 10},
	}
	for _, testcase := range testcases {
		x := testcase[0]
		y := testcase[1]
		z := testcase[2]
		expected1 := testcase[3]
		expected2 := testcase[4]
		got1, got2 := min64_2(x, y, z)
		if got1 != expected1 || got2 != expected2 {
			t.Errorf("got: (%v, %v)\nexpected: (%v, %v)", got1, got2, expected1, expected2)
		}
	}
}
