package utils

import "testing"

func TestToSnake(t *testing.T) {
	src := "sys_"
	t.Log(ToSnake(src, "_"))
}
