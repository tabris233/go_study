package intset

import (
	"math/rand"
	"testing"
	"time"
)

var n = 1000000

func BenchmarkIntSetAdd(b *testing.B) {
	var x IntSet
	rand.Seed(time.Now().Unix())
	// initData(&x, &y, xx, yy)
	for i := 0; i < b.N; i++ {
		x.Add(rand.Intn(n))
	}
}

func BenchmarkMapAdd(b *testing.B) {
	x := make(map[int]bool)
	rand.Seed(time.Now().Unix())
	// initData(&x, &y, xx, yy)
	for i := 0; i < b.N; i++ {
		x[rand.Intn(n)] = true
	}
}

// func BenchmarkUnionWith(b *testing.B) {
// var x, y IntSet
//
// }
