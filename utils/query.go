package utils

import (
	"fmt"
	"reflect"
	"strings"
)

type Field struct {
	Name  string            `json:"name"`
	Tag   reflect.StructTag `json:"tag"`
	Type  string            `json:"type"`
	Value interface{}       `json:"value"`
}

func Marshal(u interface{}) string {
	fields := []Field{}
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).CanInterface() { //判断是否为可导出字段

			//判断是否是嵌套结构
			if v.Field(i).Type().Kind() == reflect.Struct {
				structField := v.Field(i).Type()
				for j := 0; j < structField.NumField(); j++ {
					f := Field{
						Name:  structField.Field(j).Name,
						Tag:   structField.Field(j).Tag,
						Type:  structField.Field(j).Type.String(),
						Value: v.Field(i).Field(j).Interface(),
					}
					fields = append(fields, f)
				}
				continue
			}

			f := Field{
				Name:  t.Field(i).Name,
				Tag:   t.Field(i).Tag,
				Type:  t.Field(i).Type.String(),
				Value: v.Field(i).Interface(),
			}
			fields = append(fields, f)
		}
	}

	segs := []string{}
	for _, field := range fields {
		key, ok := field.Tag.Lookup("query")
		if key == "" {
			key = field.Name
		}

		val := ""
		switch field.Type {
		case "string":
			{
				val = fmt.Sprintf("%s", field.Value)
			}
		case "int", "int64":
			{
				val = fmt.Sprintf("%d", field.Value)
			}
		}

		if !strings.Contains(string(field.Tag), ",omitempty") || ok {
			segs = append(segs, key+"="+val)
		}
	}

	return strings.Join(segs, "&")
}


