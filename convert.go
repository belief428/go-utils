package utils

import (
	"reflect"
	"strings"
)

var (
	// types 数据结构类型
	types = []string{"string", "int64", "int", "uint", "uint64", "byte"}
)

// ConvertTypes 需转换的类型
type ConvertTypes struct {
	String, Int64, Int, Uint, Uint64, Byte bool
}

// TypeConvert type mutual convert
// result -- interface.(type)
func (c *ConvertTypes) TypeConvert(from interface{}) interface{} {

	fType := reflect.TypeOf(from)

	val := func(r reflect.Type) interface{} {

		return nil

	}(fType)

	return val
}

// MapToStruct map to struct
// key is struct filedName
func MapToStruct(m map[string]interface{}, s interface{}) {
	for k, v := range m {

		tRef := reflect.TypeOf(s).Elem()

		fieldNum := tRef.NumField()

		for i := 0; i < fieldNum; i++ {
			// 匹配结构字段名称
			if strings.ToLower(k) != strings.ToLower(tRef.Field(i).Name) {
				continue
			}

			vRef := reflect.ValueOf(s).Elem().FieldByName(tRef.Field(i).Name)

			if !vRef.CanSet() {
				continue
			}

			switch vRef.Type().String() {
			case "string":
				vRef.SetString(v.(string))
				break
			case "int64":
				vRef.SetInt(v.(int64))
				break
			case "int":
				switch reflect.TypeOf(v).String() {
				case "float64":
					vRef.SetInt(int64(v.(float64)))
					break
				case "int":
					vRef.SetInt(int64(v.(int)))
					break
				}
				break
			case "bool":
				vRef.SetBool(v.(bool))
				break
			}
		}
	}
}

// StructToMap struct to map
func StructToMap(s interface{}, m map[string]interface{}) {
	tRef := reflect.TypeOf(s).Elem()
	vRef := reflect.ValueOf(s).Elem()

	fieldNum := tRef.NumField()

	for index := 0; index < fieldNum; index++ {
		m[tRef.Field(index).Name] = vRef.FieldByName(tRef.Field(index).Name).Interface()
	}
}
