package utils

import (
	"fmt"
	"strconv"
)

func StringToInt(src string) int {
	dst, _ := strconv.Atoi(src)
	return dst
}

func StringToInt64(src string) int64 {
	dst, _ := strconv.ParseInt(src, 10, 64)
	return dst
}

func StringToUnit(src string) uint {
	dst, _ := strconv.Atoi(src)
	return uint(dst)
}

func StringToUnit64(src string) uint64 {
	dst, _ := strconv.ParseUint(src, 10, 64)
	return dst
}

func StringToFloat(src string) (float64, error) {
	return strconv.ParseFloat(src, 64)
}

func IntToString(src int64) string {
	return fmt.Sprintf("%d", src)
}

func UintToString(src uint64) string {
	return fmt.Sprintf("%d", src)
}
