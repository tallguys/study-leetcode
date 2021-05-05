package internal

import "math"

/**
https://leetcode-cn.com/problems/palindrome-number/
*/
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	nums := make([]int, 0)

	for {
		nums = append(nums, int(math.Mod(float64(x), 10)))
		if x = x / 10; x == 0 {
			break
		}
	}

	start := 0
	end := len(nums) - 1
	for {
		if nums[start] != nums[end] {
			return false
		}

		if end-start <= 1 {
			break
		}

		start++
		end--
	}

	return true
}
