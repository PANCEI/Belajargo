/*
value := interfaceValue.(targettipe)
*/
package main

import "fmt"

func prinvalue(i interface{}) {
	switch v := i.(type) {
	case int:
		println("Integer:", v)
	case string:
		println("String:", v)
	default:
		println("Unknown type")
	}
}
func cekvalue(i interface{}) {
	if str, ok := i.(string); ok {
		fmt.Println("String value:", str)
	} else if num, ok := i.(int); ok {
		fmt.Println("Integer value:", num)
	} else {
		fmt.Println("Unknown type")
	}
}
func main() {
	var i interface{} = "hello"
	// type assertion
	s, cek := i.(string)
	if cek {
		println(s)
	} else {
		println("Type assertion failed")
	}
	prinvalue(42)
	prinvalue("Hello, World!")
	prinvalue(3.14)
	cekvalue("hello")
	cekvalue(42)
	cekvalue(3.14)
}
