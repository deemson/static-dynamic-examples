package main

import "fmt"

func main() {
	var stringVariable string = "hello world"
	integerVariable := 42
	listOfStrings := []string{"one", "two", "three"}
	fmt.Printf("%T\n", stringVariable)
	fmt.Println(stringVariable)
	fmt.Printf("%T\n", integerVariable)
	fmt.Println(integerVariable)
	fmt.Printf("%T\n", listOfStrings)
	fmt.Println(listOfStrings)
}
