package test

import (
	"sync"
	"testing"

	"github.com/Workiva/go-datastructures/queue"
)

func BenchmarkRingBufferSPSC(b *testing.B) {
	q := queue.NewRingBuffer(uint64(b.N))
	b.ResetTimer()

	go func() {
		for i := 0; i < b.N; i++ {
			q.Put(i)
		}
	}()

	for i := 0; i < b.N; i++ {
		q.Get()
	}
}

func BenchmarkRingBufferPaddedSPSC(b *testing.B) {
	q := NewRingBufferPadded(uint64(b.N))
	b.ResetTimer()

	go func() {
		for i := 0; i < b.N; i++ {
			q.Put(i)
		}
	}()

	for i := 0; i < b.N; i++ {
		q.Get()
	}
}

func BenchmarkRingBufferMPMC(b *testing.B) {
	q := queue.NewRingBuffer(uint64(b.N * 100))
	var wg sync.WaitGroup
	wg.Add(100)
	b.ResetTimer()

	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				q.Put(i)
			}
		}()
	}

	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				q.Get()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkRingBufferPaddedMPMC(b *testing.B) {
	q := NewRingBufferPadded(uint64(b.N * 100))
	var wg sync.WaitGroup
	wg.Add(100)
	b.ResetTimer()

	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				q.Put(i)
			}
		}()
	}

	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < b.N; i++ {
				q.Get()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
