package ch12_reflect

import (
	"fmt"
	"reflect"
)

func Display(name string, v interface{}) {
	rv := reflect.ValueOf(v)
	fmt.Printf("Display %s (%s): \n", name, rv.Type())
	display(name, rv)
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			field := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(field, v.Field(i))
		}
	case reflect.Map:
		for _,key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, fmt.Sprintf("%v", key)), v.MapIndex(key))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s is nil\n", path)
		}else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s is nil\n", path)
		}else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	default:
		fmt.Printf("%s.type = %s\n", path, v.Type())
		fmt.Printf("%s.value = %v\n", path, v.Interface())
	}

}
