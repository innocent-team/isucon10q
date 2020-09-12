package main

func min64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max64(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func min64_2(x, y, z int64) (int64, int64) {
	max := max64(x, max64(y, z))
	if max == x {
		return min64(y, z), max64(y, z)
	} else if max == y {
		return min64(x, z), max64(x, z)
	}
	return min64(x, y), max64(x, y)
}
