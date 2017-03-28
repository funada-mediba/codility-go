//55%
package solution

func Solution(X int, Y int, D int) int {
	jump := X
	count := 0
	for jump < Y {
		jump = jump + D
		count++
	}
	return count
}

import "math"
func Solution(X int, Y int, D int) int {
	Xfloat64 := float64(X)
	Yfloat64 := float64(Y)
	Dfloat64 := float64(D)
	buffer := ((Yfloat64 - (Xfloat64 + Dfloat64)) / Dfloat64)
	result := math.Floor(buffer + 1.5)
	return int(result)
}
