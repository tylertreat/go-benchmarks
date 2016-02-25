package test

import (
	"sync"
	"testing"
)

func BenchmarkMutexDeferUnlock(b *testing.B) {
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		func() {
			mu.Lock()
			defer mu.Unlock()
		}()
	}
}

func BenchmarkMutexUnlock(b *testing.B) {
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		func() {
			mu.Lock()
			mu.Unlock()
		}()
	}
}
