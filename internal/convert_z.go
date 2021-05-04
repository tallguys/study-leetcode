package internal

import (
	"fmt"
	"strings"
)

/**
https://leetcode-cn.com/problems/zigzag-conversion/
*/
func ConvertZ(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	res := make([]string, numRows)
	for i := range res {
		res[i] = ""
	}

	flag := -1
	i := 0

	for _, r := range s {
		res[i] = res[i] + fmt.Sprintf("%c", r)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}

	return strings.Join(res, "")
}
