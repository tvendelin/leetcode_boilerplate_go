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
		if false { // replace with actual test
			t.Errorf("%v expected %d, got %d\n", tc.I, tc.E, actual)
		}
	}
}

/* uncomment for tasks with linked lists
func toLinkedList(s []int) *ListNode {
	var last *ListNode
	var node *ListNode
	for i := len(s) - 1; i >= 0; i-- {
		node = &ListNode{Val: s[i], Next: last}
		last = node
	}
	return node
}
*/
