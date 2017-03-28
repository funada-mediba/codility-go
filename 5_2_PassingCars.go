//50%
func Solution(A []int) int {
	count := 0
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			for j := i; j < len(A); j++ {
				if A[j] == 1 {
					count++
				}
			}
		}
	}
	return count
}
