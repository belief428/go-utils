package utils

import (
	"fmt"
	"reflect"
)

func InArray(search, needle interface{}) bool {
	val := reflect.ValueOf(needle)

	kind := val.Kind()

	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < val.Len(); i++ {
			if val.Index(i).Interface() == search {
				return true
			}
		}
	}
	return false
}

func ArrayFlip(src interface{}) interface{} {
	val := reflect.ValueOf(src)

	kind := val.Kind()

	out := make(map[interface{}]interface{}, 0)

	if kind == reflect.Slice || kind == reflect.Array || kind == reflect.Map {
		for i := 0; i < val.Len(); i++ {
			fmt.Printf("val.Field()：%v\n", val.Index(i).MapKeys())
		}
	}
	return out
}

func ArrayStrings(src interface{}) []string {
	out := make([]string, 0)

	val := reflect.ValueOf(src)

	kind := val.Kind()

	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < val.Len(); i++ {
			out = append(out, fmt.Sprintf("%v", val.Index(i).Interface()))
		}
	}
	return out
}

func ArrayUnique(src interface{}) []interface{} {
	val := reflect.ValueOf(src)

	kind := val.Kind()

	out := make([]interface{}, 0)

	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < val.Len(); i++ {
			for _, v := range out {
				if v == val.Index(i).Interface() {
					goto BREAK
				}
			}
			out = append(out, val.Index(i).Interface())
		BREAK:
		}
	}
	return out
}
