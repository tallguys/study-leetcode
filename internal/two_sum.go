package internal

/**
https://leetcode-cn.com/problems/two-sum/
*/
func TwoSum(nums []int, target int) []int {
	store := make(map[int]int)
	for i, num := range nums {
		if j, ok := store[target-num]; ok {
			return []int{i, j}
		}

		store[num] = i
	}

	return nil
}
