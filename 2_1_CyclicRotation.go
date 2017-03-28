//100%
package solution

func Solution(A []int, K int) []int {
	if 1 < len(A) {
		for i := 0; i < K; i++ {
			buffer := A[len(A)-1]
			A = A[:len(A)-1]
			A, A[0] = append(A[0:1], A[0:]...), buffer
		}
	}
	return A
}
