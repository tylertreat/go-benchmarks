package test

import (
	"strconv"
	"sync"
	"testing"

	"github.com/Workiva/go-datastructures/trie/ctrie"
)

func snapshotMap(m map[string]struct{}, mu *sync.Mutex) map[string]struct{} {
	mu.Lock()
	snapshot := make(map[string]struct{}, len(m))
	for k, v := range m {
		snapshot[k] = v
	}
	mu.Unlock()
	return snapshot
}

func BenchmarkSnapshotMap(b *testing.B) {
	m := make(map[string]struct{})
	for i := 0; i < b.N; i++ {
		m[strconv.Itoa(i)] = struct{}{}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		snapshot := make(map[string]struct{}, b.N)
		for k, v := range m {
			snapshot[k] = v
		}
	}
}

func BenchmarkSnapshotCtrie(b *testing.B) {
	m := ctrie.New(nil)
	for i := 0; i < b.N; i++ {
		m.Insert([]byte(strconv.Itoa(i)), struct{}{})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Snapshot()
	}
}

func BenchmarkConcurrentSnapshotMap(b *testing.B) {
	m := make(map[string]struct{})
	for i := 0; i < b.N; i++ {
		m[strconv.Itoa(i)] = struct{}{}
	}
	mu := &sync.Mutex{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(100)
		for j := 0; j < 100; j++ {
			go func() {
				snapshotMap(m, mu)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}

func BenchmarkConcurrentSnapshotCtrie(b *testing.B) {
	m := ctrie.New(nil)
	for i := 0; i < b.N; i++ {
		m.Insert([]byte(strconv.Itoa(i)), struct{}{})
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		wg.Add(100)
		for j := 0; j < 100; j++ {
			go func() {
				m.Snapshot()
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
