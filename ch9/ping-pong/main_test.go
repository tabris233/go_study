package main

import "testing"

func Benchmarkch1ch2(b *testing.B) {
	done := make(chan struct{})
	ch1 := make(chan int)
	ch2 := make(chan int)

	// TODO 又不死锁了。。。。。
	go func() {
		for i := 0; i < b.N; i++ {
			ch1 <- 1
			<-ch2
		}
		done <- struct{}{}
	}()
	go func() {
		for i := 0; i < b.N; i++ {
			ch2 <- 2
			<-ch1
		}
		done <- struct{}{}
	}()

	<-done
	<-done
	close(done)
	// for range ch1 {
	// }
	// for range ch2 {
	// }
	close(ch1)
	close(ch2)
}
