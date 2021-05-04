package internal

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	formatReg = regexp.MustCompile("^[ ]*[+-]?[0-9]+")
	numReg    = regexp.MustCompile("^[0-9]$")
)

const (
	_space    = " "
	_negative = "-"
	_positive = "+"
)

/**
https://leetcode-cn.com/problems/string-to-integer-atoi/
*/
func MyAtoi(s string) int {
	if !formatReg.MatchString(s) {
		return 0
	}

	s = strings.TrimLeft(s, _space)

	numStr := ""
	start := 0

	if s[0:1] == _positive || s[0:1] == _negative {
		numStr += s[0:1]
		start = 1
	}

	for _, r := range s[start:] {
		if numReg.MatchString(string(r)) {
			numStr += string(r)
		} else {
			break
		}
	}

	num, err := strconv.Atoi(numStr)
	if err != nil {
		if num > math.MaxInt32 {
			return math.MaxInt32
		} else if num < math.MinInt32 {
			return math.MinInt32
		}

		return 0
	}

	if num > math.MaxInt32 {
		return math.MaxInt32
	} else if num < math.MinInt32 {
		return math.MinInt32
	}

	return num
}
