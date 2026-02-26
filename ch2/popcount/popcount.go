package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the number of set bits of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Exercise 2.3: Rewrite PopCount to use a loop instead of a single expression.
// Compare the performance of the two versions.
func Exercise_2_3(x uint64) int {
	result := byte(0)
	for i := range 8 {
		result += pc[byte(x>>(i*8))]
	}
	return int(result)
}

// Exercise 2.4: Write a version of PopCount that counts bits by shifting its
// argument through 64 bit positions, testing the rightmost bit each time.
// Compare its performance to the table-lookup version.
func Exercise_2_4(x uint64) int {
	result := 0
	for i := range 64 {
		result += int((x >> i) & 1)
	}
	return result
}

// Exercise 2.5: The expression x&(x-1) clears the right most non-zero bit of x.
// Write a version of PopCount that counts bits by using this fact, and assess
// its performance.
func Exercise_2_5(x uint64) int {
	result := 0
	for x != 0 {
		x = x & (x - 1)
		result++
	}
	return result
}

func TestPerformance() int {
	sum := 0
	for i := range 10000000 {
		sum += Exercise_2_5(uint64(i))
	}
	return sum
}
