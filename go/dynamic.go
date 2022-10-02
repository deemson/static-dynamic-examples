package main

import (
	"fmt"
	"strings"
)

func main() {
	var dynamicVariable interface{}
	dynamicVariable = "hello"
	fmt.Printf("%T\n", dynamicVariable)
	fmt.Println(dynamicVariable)
	fmt.Println(strings.ToUpper(dynamicVariable.(string)))
	dynamicVariable = 42
	fmt.Printf("%T\n", dynamicVariable)
	fmt.Println(dynamicVariable)
	fmt.Println(strings.ToUpper(dynamicVariable.(string)))
}
