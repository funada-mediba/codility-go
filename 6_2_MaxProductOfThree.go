//55%
package solution

import "sort"

func Solution(A []int) int {
	result := 0
	sort.Sort(sort.IntSlice(A))

	result = A[len(A)-1] * A[len(A)-2]
	if len(A) == 3 {
		result = A[len(A)-1] * A[len(A)-2] * A[len(A)-3]
	} else {
		if A[len(A)-3] == A[len(A)-4] {
			result = A[len(A)-1] * A[len(A)-3] * A[len(A)-4]
		} else {
			result = result * A[len(A)-3]
		}
	}
	return result
}
