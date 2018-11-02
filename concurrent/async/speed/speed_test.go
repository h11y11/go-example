package speed_test

import (
	"testing"

	"github.com/iashbringer/go-example/basework"
)

func BenchmarkGoroutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go basework.SimpleWork()
	}
}

func benchmarkChan(b *testing.B, buflen int) {
	emptyObj := struct{}{}
	ch := make(chan struct{}, buflen)
	go func() {
		for range ch {
			basework.SimpleWork()
		}
	}()
	for i := 0; i < b.N; i++ {
		ch <- emptyObj
	}
	close(ch)
}

func BenchmarkChanNoBuffer(b *testing.B) {
	benchmarkChan(b, 0)
}
func BenchmarkChanBuffer100(b *testing.B) {
	benchmarkChan(b, 100)
}

func BenchmarkChanBuffer1000(b *testing.B) {
	benchmarkChan(b, 1000)
}

func BenchmarkChanBuffer10000(b *testing.B) {
	benchmarkChan(b, 10000)
}
