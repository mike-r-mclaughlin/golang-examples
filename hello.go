package main

import "fmt"

func main() {
	fmt.Println("hello world")
	var a_number = 4321
	var a_number_as_str = fmt.Sprintf("%d", a_number)
	a_number_as_str = fmt.Sprintf ("%d", a_number)
	a_number_as_str = fmt.Sprintf( "%d", a_number)
	a_number_as_str = fmt.Sprintf(
		"%d", a_number)
	fmt.Printf("The value is %s\n", a_number_as_str)
	a_number_as_str = fmt.Sprintf(
		// convert the number + 10 to a string
		"%d", a_number + 10)
	fmt.Printf("The new value is %s\n", a_number_as_str)
}

