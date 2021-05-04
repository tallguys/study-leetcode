package internal

import "math"

/**
https://leetcode-cn.com/problems/reverse-integer/
*/
func Reverse(x int) int {
	ret := 0

	for {
		m := math.Mod(float64(x), float64(10))
		if x/10 == 0 && m == 0 {
			break
		}
		x = x / 10
		ret = ret*10 + int(m)
	}

	if ret > math.MaxInt32 || ret < math.MinInt32 {
		return 0
	}

	return ret
}
