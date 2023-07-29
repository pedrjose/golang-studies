package main

import "fmt"

func main() {
	ex8()
}

// #1 Print a hello world using go
func ex1() {
	fmt.Println(`"Hello, World"`)
}

// #2 Create 2 variables, x and y, a string and a int, and print the value and the type
func ex2() {
	x := 10
	y := "Good Morning!\nHow are you?\tToday is \"friday\"."

	fmt.Printf("Value x: %v,| Type x: %T\n", x, x)
	fmt.Printf("Value y: %v,| Type y: %T", y, y)
}

// #3 Create the type variable "wage" and print it
func ex3() {
	type wage int
	var pay wage = 1250

	fmt.Printf("My salary is %v and it' declared like %T", pay, pay)
}

// #4 Convert a main value to a "fake" type and print the before and the after
func ex4() {
	type FakeType float64

	mainValue := 42.00
	fmt.Printf("Before Conversion: %v (Type: '%T')\n\n", mainValue, mainValue)

	fakeNumber := FakeType(mainValue)
	fmt.Printf("Fake Number After Conversion: %v (Type: '%T')", fakeNumber, fakeNumber)
}

/* #5 Using the short declaration operator (Gopher), assign these values to variables with the identifiers "x", "y", and "z": 42, "James Bond", true. Now, demonstrate the values in these variables with: A single print statement and Multiple print statements */
func ex5() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("\n\n'x' value: %v | 'y' value: %v | 'z' value: %v\n\n", x, y, z)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

/* #6 Use "var" to declare three variables. They must have package-level scope. Do not assign values to these variables. Use the following identifiers and types for these variables:

Identifier "x" should have type int
Identifier "y" should have type string
Identifier "z" should have type bool

In the "main" function:
- 1. Display the values of each identifier
- 2. The compiler assigned values to these variables. What are these values called? */
var x int
var y string
var z bool

func ex6() {
	fmt.Printf("\n\n1 - 'x' value: %v | 'y' value: %v | 'z' value: %v\n", x, y, z)
	fmt.Println("2 - When variables are not initialized, the compiler assigned with the '0' value")
}

/* #7 Create a type. The underlying type should be int. Create a variable for this type, with the identifier "q", using the keyword var. In the "main" function:

1 - Display the value of the variable "x".
2 - Display the type of the variable "x".
3 - Assign 42 to the variable "x" using the "=" operator.
4 - Display the value of the variable "x". */
type fakeNumber int

var q fakeNumber

func ex7() {
	fmt.Printf("1 - Value of 'q': %v\n", q)
	fmt.Printf("2 - Type of 'q': %T\n", q)

	// 3
	x = 42
	fmt.Println("3 - 'q = 42'")
	fmt.Printf("4 - Value of 'q' after assigned 42: %v", q)
}

/* #8 Using the solution from the previous exercise: In package-level scope, using the keyword var, create a variable with the identifier "u". The type of this variable should be the underlying type of the type you created in the previous exercise. In the previous exercise, this should already be done:

1 - Display the value of the variable "q".
2 - Display the type of the variable "q".
3 - Assign 42 to the variable "q" using the "=" operator.
4 - Display the value of the variable "q".

Now, do the following as well:

1 - Use conversion to transform the type of the value of variable "x" into its underlying type, and using the "=" operator, assign the value of "q" to "u".
2 - Display the value of "u".
3 - Display the type of "u". */
var u int

func ex8() {
	u = int(q)
	fmt.Printf("'u' value: %v\n", u)
	fmt.Printf("'u' type: %T", u)
}
