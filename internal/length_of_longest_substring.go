package internal

/**
https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
*/
func LengthOfLongestSubstring(s string) int {
	maxLen := 0
	tmpSet := make(map[rune]int)
	start := 0

	for end, r := range s {
		if index, ok := tmpSet[r]; ok && index >= start {
			start = index + 1
		}

		tmpSet[r] = end

		if end-start+1 > maxLen {
			maxLen = end - start + 1
		}
	}
	return maxLen
}
