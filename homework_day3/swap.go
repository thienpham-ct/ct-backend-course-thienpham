// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Swap
//
//  Using funcs, swap values through pointers, and swap
//  the addresses of the pointers.
//
//
//  1- Swap the values through a func
//
//     a- Declare two float variables
//
//     b- Declare a func that can swap the variables' values
//        through their memory addresses
//
//     c- Pass the variables' addresses to the func
//
//     d- Print the variables
//
//  2- Swap the addresses of the float pointers through a func
//
//     a- Declare two float pointer variables and,
//        assign them the addresses of float variables
//
//     b- Declare a func that can swap the addresses
//        of two float pointers
//
//     c- Pass the pointer variables to the func
//
//     d- Print the addresses, and values that are
//        pointed by the pointer variables
//
// ---------------------------------------------------------

package main

import "fmt"

func swapValues(a, b float32) (float32, float32) {
	return b, a
}

func swapAddress(ap, bp *float32) {
	*ap, *bp = *bp, *ap
}

func main() {
	a := float32(2)
	b := float32(4)

	fmt.Println("Pointer address of a", &a, "value of a", a)
	fmt.Println("Pointer address of b", &b, "value of b", b)

	a, b = b, a

	fmt.Println("Pointer address of a after value swap", &a, "value of a", a)
	fmt.Println("Pointer address of b after value swap", &b, "value of b", b)

	var ap *float32
	var bp *float32
	ap = &a
	bp = &b

	swapAddress(ap, bp)

	fmt.Println("Pointer address of a after address swap", ap, "value of a", *ap)
	fmt.Println("Pointer address of b after address swap", bp, "value of b", *bp)
}
