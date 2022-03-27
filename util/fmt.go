package util

import "fmt"

func Brk() {
	fmt.Println("============================")
}

func PrintConcreteType(v interface{}) {
	fmt.Printf("{type : \"%T\", value: \"%#v\"}\n", v, v)
}
