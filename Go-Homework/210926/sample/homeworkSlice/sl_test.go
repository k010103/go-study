package homeworkSlice

import "testing"

// b.N 반복문을 돌려야하는 이유
// 같은 작업을 여러번 동작하면서 프로그램 상에서 평균값을 내어준다.
// 그래서 속도를 테스트 할 경우 상수를 넣어서 동일한 작업을 반복하는 것이 맞다.

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
