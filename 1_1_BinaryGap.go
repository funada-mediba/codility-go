//100%すごい
package solution

func Solution(N int) int {
	var bit []int
	for i := N; i != 0; {
		bit = append(bit, i%2)
		i = i / 2
	}
	flag := false
	count := 0
	result := 0
	for i := 0; i < len(bit); i++ {
		if bit[i] == 1 {
			flag = true
		}
		if bit[i] == 0 && flag == true {
			count++
		}
		if bit[i] == 1 && flag == true {
			if result < count {
				result = count
			}
			count = 0
		}
	}
	return result
}
