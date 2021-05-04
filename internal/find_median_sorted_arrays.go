package internal

import "math"

/**
https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
*/
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	cur := 0
	last := 0
	i := 0
	p1 := 0
	p2 := 0

	for {
		if p1 < len(nums1) && p2 < len(nums2) {
			if nums1[p1] < nums2[p2] {
				cur = nums1[p1]
				p1++
			} else {
				cur = nums2[p2]
				p2++
			}
		} else if p1 < len(nums1) && p2 >= len(nums2) {
			cur = nums1[p1]
			p1++
		} else if p1 >= len(nums1) && p2 < len(nums2) {
			cur = nums2[p2]
			p2++
		}

		count := len(nums1) + len(nums2)
		if math.Mod(float64(count), 2) == 0 {
			if i == count/2 {
				return (float64(last) + float64(cur)) / 2
			}
		} else {
			if i == count/2 {
				return float64(cur)
			}
		}

		i++
		last = cur
	}
}
