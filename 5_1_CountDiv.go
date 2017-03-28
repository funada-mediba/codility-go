//50%
func Solution(A int, B int, K int) int {
	count := 0
	for A <= B {
		result := A % K
		if result == 0 {
			count++
		}
		A++
	}
	return count
}
