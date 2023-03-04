package main

import "fmt"

var k int
var l int = 2
var m = 3

// z := 2 error can onlu be used inside functions
// Variable declaration and value assignment cannot be done separately (must be done in the same line)

func main() {
	// fmt.Println(z)
	//using var can be used inside and outside of functions
	//variable declaration and value assignment can be done separately
	k = 1
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(m)

	//Declaring(Creating) Variables
	var student1 string = "John" // with the var keyword
	var student2 = "Jane"        //type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	//Variable Declaration Without Initial Value
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	//Variable Assignment After Declaration
	var student3 string
	student3 = "Johnny"
	fmt.Println(student3)

	//Go Multiple Variable Declaration
	var q, w, e, r int = 1, 3, 5, 7

	fmt.Println(q)
	fmt.Println(w)
	fmt.Println(e)
	fmt.Println(r)

	//if the type keyword is not specified, you can declare different types of variables int the same line:
	var nomer, kata = 6, "Hello"
	nomer1, kata1 := 7, "World!"

	fmt.Println(nomer)
	fmt.Println(kata)
	fmt.Println(nomer1)
	fmt.Println(kata1)

	//Go Variable Declaration in a Block
	var (
		null int
		satu int    = 1
		word string = "hello"
	)

	fmt.Println(null)
	fmt.Println(satu)
	fmt.Println(word)
}

/*Go Variable Naming Rules
A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

Go variable naming rules:

A variable name must start with a letter or an underscore character (_)
A variable name cannot start with a digit
A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
Variable names are case-sensitive (age, Age and AGE are three different variables)
There is no limit on the length of the variable name
A variable name cannot contain spaces
The variable name cannot be any Go keywords*/

// Camel Case
/* myVariableName = "John"*/

// Pascal Case
/* MyVariableName = "John"*/

// Snake Case
/*my_variable_name = "John"*/
