package stringutil

import (
	"github.com/zonewave/pkgs/standutil/expr"
	"github.com/zonewave/pkgs/standutil/randutil"
	"math/rand"
	"strings"
	"unicode"
)

// RandString returns a string representation of random
func RandString(lenNum int, rands ...*rand.Rand) string {
	r := expr.FirstOrDefault(rands, randutil.DefaultRand())
	str := strings.Builder{}
	length := 52
	for i := 0; i < lenNum; i++ {
		str.WriteString(CHARS[r.Intn(length)])
	}
	return str.String()
}

// SpaceRemove remove n space
func SpaceRemove(s string, n int) string {
	newStr := ""
	if n == 0 {
		return s
	}
	for _, c := range s {
		if unicode.IsSpace(c) {
			if n == -1 {
				continue
			}
			if n > 0 {
				n--
				continue
			}
		}
		newStr += string(c)
	}
	return newStr
}

// SpaceRemoveAll remove all space
func SpaceRemoveAll(s string) string {
	return SpaceRemove(s, -1)
}
