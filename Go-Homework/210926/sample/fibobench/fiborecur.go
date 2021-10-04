package fibobench

func fiboRecur(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 || n == 1 {
		return 1
	}
	return fiboRecur(n-1) + fiboRecur(n-2)
}
