package benchmark

import "testing"

func BenchmarkSelect(b *testing.B) {
	InitServerIndex()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FastRandSelect()
	}
}

func BenchmarkSelectParallel(b *testing.B) {
	InitServerIndex()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			FastRandSelect()
		}
	})
}
