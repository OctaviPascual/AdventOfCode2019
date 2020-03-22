package util

// Abs returns the absolute value on n
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// GCD returns the greatest common divisor of a and b
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// LCM returns the least common multiple of a and b
func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}

// LCM3 returns the least common multiple of a, b and c
func LCM3(a, b, c int) int {
	return LCM(LCM(a, b), c)
}
