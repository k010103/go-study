package fibobench

import "testing"

func BenchmarkFiboDP(b *testing.B) {
	for i := 0 ; i < b.N ; i ++ {
		fiboDP(i)
	}
}

func BenchmarkFiboRecur(b *testing.B) {
	for i := 0 ; i < b.N ; i ++ {
		fiboRecur(i)
	}
}
