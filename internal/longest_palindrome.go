package internal

/**
https://leetcode-cn.com/problems/longest-palindromic-substring/
*/
func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	ret := s[0:1]

	for i := range s {
		p1 := palindrome(s, i, i)
		if len(p1) > len(ret) {
			ret = p1
		}

		if i < len(s)-1 {
			p2 := palindrome(s, i, i+1)
			if len(p2) > len(ret) {
				ret = p2
			}
		}
	}

	return ret
}

func palindrome(s string, start int, end int) string {
	ret := s[start : start+1]

	for {
		if start < 0 || end > len(s)-1 {
			break
		}

		if s[start] != s[end] {
			break
		}

		ret = s[start : end+1]

		start--
		end++
	}

	return ret
}
