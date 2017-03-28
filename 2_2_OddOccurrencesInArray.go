//100%
package solution

import "sort"

func Solution(A []int) int {
	sort.Sort(sort.IntSlice(A))
	if len(A) == 1 {
		return A[0]
	}
	result := 0
	for i := 0; i < len(A); i++ {
		if i%2 != 0 {
			if A[i] != A[i-1] {
				result = A[i-1]
				break
			}
		} else {
			result = A[i]
		}
	}
	return result
}
