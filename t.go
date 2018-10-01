package main

import "github.com/allbuleyu/gobase/gopl.io/ch12-reflect"

func main() {
	i := 123
	ch12_reflect.Display("i", &i)

	m := []map[string]string{
		{
			"sss" : "sss",
			"dddd" : "dddd",
		},
		{
			"sss" : "sss",
			"dddd" : "dddd",
		},
		{
			"sss" : "sss",
			"dddd" : "dddd",
		},
	}

	ch12_reflect.Display("M", m)
}
