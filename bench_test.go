package main

import "testing"

func BenchmarkExtraFor(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		// benchmark target
	}
}
