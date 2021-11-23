package utils

import (
	"strings"
)

func ToLowerCase(src []rune) rune {
	return src[0] | 0x20
}

func ToUpperCase(src []rune) rune {
	return src[0] ^ 0x20
}

func ToSnake(src, delimiter string) string {
	src = strings.TrimSpace(src)
	objs := strings.Split(src, delimiter)

	out := make([]rune, 0)

	for _, v := range objs {
		if len(v) <= 0 {
			continue
		}
		obj := []rune(v)
		obj[0] = ToUpperCase(obj)
		out = append(out, obj...)
	}
	return string(out)
}
