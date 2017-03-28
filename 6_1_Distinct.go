//100%
import "sort"

func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}
	sort.Sort(sort.IntSlice(A))
	buff := A[0]
	count := 1

	for i := range A {
		if buff != A[i] {
			count++
			buff = A[i]
		}
	}
	return count
}
