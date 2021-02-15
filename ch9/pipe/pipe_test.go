package main

import "testing"

func BenchmarkPipeLime(b *testing.B) {
	in, out := pipeline(1000000)
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
	close(in)
}
