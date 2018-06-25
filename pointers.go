package main

import "fmt"

func main() {
	x := 15
	a := &x // memory address
	fmt.Println("Address:", a)
	fmt.Println("Value from address:", *a) // dereference address

	*a = 5 // change value based off address
	fmt.Println("New Val:", x)
}
