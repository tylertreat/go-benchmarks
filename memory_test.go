package test

import (
	"sync"
	"testing"
)

func BenchmarkAllocateHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() *Struct {
			return new(Struct)
		}()
	}
}

func BenchmarkAllocateStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() Struct {
			return Struct{}
		}()
	}
}

func BenchmarkPassByReference(b *testing.B) {
	f := func(s *Struct) {}
	s := new(Struct)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(s)
	}
}

func BenchmarkPassByValue(b *testing.B) {
	f := func(s Struct) {}
	s := Struct{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		f(s)
	}
}

func BenchmarkConcurrentStructAllocate(b *testing.B) {
	var wg sync.WaitGroup
	wg.Add(10)
	b.ResetTimer()

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < b.N; j++ {
				_ = func() *Struct {
					return new(Struct)
				}()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

func BenchmarkConcurrentStructPool(b *testing.B) {
	pool := &sync.Pool{New: func() interface{} { return new(Struct) }}
	var wg sync.WaitGroup
	wg.Add(10)
	b.ResetTimer()

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < b.N; j++ {
				s := func() *Struct {
					return pool.Get().(*Struct)
				}()
				pool.Put(s)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
