package main

import (
	"testing"
)

type testcase struct {
	I []int
	E int
}

var cases []testcase = []testcase{}

func TestT(t *testing.T) {
	for _, tc := range cases {
		actual := 0
		if true { // replace with actual test
			t.Errorf("%v expected %d, got %d\n", tc.I, tc.E, actual)
		}
	}
}
