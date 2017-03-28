//100%
import "sort"

func Solution(A []int) int {
	sort.Sort(sort.IntSlice(A))
	for i, v := range A {
		if i+1 != v {
			return 0
			break
		}
	}
	return 1
}
