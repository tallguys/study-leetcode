package internal

/**
https://leetcode-cn.com/problems/binary-search/
*/
func Search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for {
		if left > right {
			return -1
		}

		mid := (right + left) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
}
