package homeworkSlice

import "testing"

func BenchmarkTenHun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slice(10, 100)
	}
}

func BenchmarkHunHun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slice(100, 100)
	}
}
func BenchmarkTenTho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slice(10, 1000)
	}
}
func BenchmarkHunTho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slice(100, 1000)
	}
}
func BenchmarkThoTho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Slice(1000, 1000)
	}
}

// func BenchmarkSlice(b *testing.B) {
// 	Slice(10, 100)
// 	Slice(100, 100)
// 	Slice(10, 1000)
// 	Slice(100, 1000)
// 	Slice(1000, 1000)
// }
