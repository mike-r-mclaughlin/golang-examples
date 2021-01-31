package main

import "fmt"

func main() {
	fmt.Println("hello world")
	var a_number = 4321
	var a_number_as_str = fmt.Sprintf(
		"%d", a_number)
	fmt.Printf("The value is %s\n", a_number_as_str)
}

