package test

import (
	"sync"
	"testing"

	"github.com/Workiva/go-datastructures/queue"
)

func BenchmarkChannel(b *testing.B) {
	ch := make(chan interface{}, 1)

	b.ResetTimer()
	go func() {
		for i := 0; i < b.N; i++ {
			<-ch
		}
	}()

	for i := 0; i < b.N; i++ {
		ch <- `a`
	}
}

func BenchmarkRingBuffer(b *testing.B) {
	q := queue.NewRingBuffer(1)

	b.ResetTimer()
	go func() {
		for i := 0; i < b.N; i++ {
			q.Get()
		}
	}()

	for i := 0; i < b.N; i++ {
		q.Put(`a`)
	}
}

func BenchmarkChannelReadContention(b *testing.B) {
	ch := make(chan interface{}, 100)
	var wg sync.WaitGroup
	wg.Add(1000)
	b.ResetTimer()

	go func() {
		for i := 0; i < b.N; i++ {
			ch <- `a`
		}
	}()

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < b.N/1000; i++ {
				<-ch
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkRingBufferReadContention(b *testing.B) {
	q := queue.NewRingBuffer(100)
	var wg sync.WaitGroup
	wg.Add(1000)
	b.ResetTimer()

	go func() {
		for i := 0; i < b.N; i++ {
			q.Put(`a`)
		}
	}()

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < b.N/1000; i++ {
				q.Get()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkChannelContention(b *testing.B) {
	ch := make(chan interface{}, 100)
	var wg sync.WaitGroup
	wg.Add(1000)
	b.ResetTimer()

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				ch <- `a`
			}
		}()
	}

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				<-ch
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkRingBufferContention(b *testing.B) {
	q := queue.NewRingBuffer(100)
	var wg sync.WaitGroup
	wg.Add(1000)
	b.ResetTimer()

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				q.Put(`a`)
			}
		}()
	}

	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				q.Get()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
