package mstrings

import "unicode"

func SpaceRemove(s string, n int) string {
	newStr := ""
	if n == 0 {
		return s
	}
	for _, c := range s {
		if n == -1 || n > 0 {
			if unicode.IsSpace(c) {
				if n != -1 {
					n--
				}
				continue
			}
		}
		newStr += string(c)
	}
	return newStr
}

func SpaceRemoveAll(s string) string {
	return SpaceRemove(s, -1)
}
