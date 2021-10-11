package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	ss := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		s := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			s[j] = uint8(i * j)
		}
		ss[i] = s
	}
	return ss
}

func main() {
	pic.Show(Pic)
}
