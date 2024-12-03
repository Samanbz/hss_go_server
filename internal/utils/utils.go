package utils

import (
	"fmt"
	"reflect"
)

func InspectStruct(ptr interface{}) {
	v := reflect.ValueOf(ptr)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := v.Type().Field(i)

		fmt.Printf("%s (%s): ", fieldType.Name, fieldType.Type)

		if field.Kind() == reflect.Ptr {

			if field.IsNil() {
				fmt.Println("nil")
			} else if field.Elem().Kind() == reflect.Array {
				for i := 0; i < field.Elem().Len(); i++ {
					InspectStruct(field.Elem().Index(i).Addr().Interface())
				}
			} else {
				fmt.Println(reflect.Indirect(field).Interface())
			}
		} else {
			fmt.Println(field.Interface())
		}
	}

}
