package main

import "fmt"

func main() {
	ex4()
}

// Print a hello world using go
func ex1() {
	fmt.Println(`"Hello, World"`)
}

// Create 2 variables, x and y, a string and a int, and print the value and the type
func ex2() {
	x := 10
	y := "Good Morning!\nHow are you?\tToday is \"friday\"."

	fmt.Printf("Value x: %v,| Type x: %T\n", x, x)
	fmt.Printf("Value y: %v,| Type y: %T", y, y)
}

// Create the type variable "wage" and print it
func ex3() {
	type wage int
	var pay wage = 1250

	fmt.Printf("My salary is %v and it' declared like %T", pay, pay)
}

// Convert a main value to a "fake" type and print the before and the after
func ex4() {
	type FakeType float64

	mainValue := 42.00
	fmt.Printf("Before Conversion: %v (Type: '%T')\n\n", mainValue, mainValue)

	fakeNumber := FakeType(mainValue)
	fmt.Printf("Fake Number After Conversion: %v (Type: '%T')", fakeNumber, fakeNumber)
}
