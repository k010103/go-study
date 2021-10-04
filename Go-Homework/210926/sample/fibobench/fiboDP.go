package fibobench

func fiboDP(n int) int {
	if n < 0 {
		return -1
	}
	if n == 0 || n == 1 {
		return 1
	}
	table := make([]int, n)
	table[0], table[1] = 1, 1
	for i := 2 ; i < n ; i ++ {
		table[i] = table[i-1] + table[i-2]
	}
	return table[n-1]
}
