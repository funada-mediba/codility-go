//100%
import "sort"

func Solution(A []int) int {
	sort.Sort(sort.IntSlice(A))
	if len(A) == 0 {
		return 1
	}
	result := 0
	for i, v := range A {
		if i+1 != v {
			result = i + 1
			break
		} else {
			result = v + 1
		}
	}
	return result
}
