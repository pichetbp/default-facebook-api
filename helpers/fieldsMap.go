package helpers

import "reflect"

func GetFieldMap[T any](obj T) string {
	resp := ""
	v := reflect.ValueOf(obj)
	for i := 0; i < v.Type().NumField(); i++ {
		if i > 0 {
			resp += ","
		}
		resp += v.Type().Field(i).Tag.Get("json")
	}
	return resp
}
