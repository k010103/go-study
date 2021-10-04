package homeworkSlice

func Slice(c, l int) []int {
	a := make([]int, 0, c)

	for i := 0; i < l; i++ {
		a = append(a, i)
	}
	return a
	// fmt.Printf("len=%d cap=%d\n", len(a), cap(a))
}
