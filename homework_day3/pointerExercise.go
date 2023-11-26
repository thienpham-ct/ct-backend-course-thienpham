// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Basics
//
//  Let's warm up with the pointer basics. Please follow the
//  instructions inside the code. Run the solution to see
//  its output if you need to.
// ---------------------------------------------------------

package main

import "fmt"

type computer struct {
	brand string
}

func main() {
	// create a nil pointer with the type of pointer to a computer
	var apple_p *computer
	// compare the pointer variable to a nil value, and say it's nil
	if apple_p == nil {
		fmt.Println("It's nil")
	}
	// create an apple computer by putting its address to a pointer variable
	apple_p = &computer{"Apple"}

	// put the apple into a new pointer variable
	newp := apple_p
	// compare the apples: if they are equal say so and print their addresses
	if apple_p == newp {
		fmt.Println("equal. Address is", &apple_p)
	}
	// create a sony computer by putting its address to a new pointer variable
	sony_p := &computer{"Sony"}

	// compare apple to sony, if they are not equal say so and print their
	// addresses
	if sony_p != apple_p {
		fmt.Println("not equal. Address of sony:", &sony_p, ". Address of apple", &apple_p)
	}

	// put apple's value into a new ordinary variable
	nov := *apple_p
	// print apple pointer variable's address, and the address it points to
	// and, print the new variable's addresses as well
	fmt.Println("Apple pointer address:", &apple_p, "Address points to: ", apple_p)
	fmt.Println("New ordinary variable address:", &nov)

	// compare the value that is pointed by the apple and the new variable
	// if they are equal say so
	if *apple_p == nov {
		fmt.Println("Value of apple_p and nov is equal")
	}
	// print the values:
	// the value that is pointed by the apple and the new variable
	fmt.Println("Value of apple_p", *apple_p)
	fmt.Println("Value of nov", nov)

	// change sony's brand to hp using the func — print sony's brand
	change(sony_p, "HP")
	fmt.Println("Sony's brand:", *sony_p)

	// call the constructor 3 times and print the returned values' addresses
	apple := computer{"Apple"}
	fmt.Println("Call the constructor 3 times")
	fmt.Println(constructor(apple))
	fmt.Println(constructor(apple))
	fmt.Println(constructor(apple))

}

// create a new function: change
// the func can change the given computer's brand to another brand
func change(c *computer, newBrand string) {
	c.brand = newBrand
}

// write a func that returns the value that is pointed by the given *computer
// print the returned value
func printValue(c *computer) {
	fmt.Println(*c)
}

// write a new constructor func that returns a pointer to a computer
// and call the func 3 times and print the returned values' addresses
func constructor(c computer) *computer {
	return &c
}
