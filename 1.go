package main

import "fmt"

// Comments

/* Multiline */

func main() {

  fmt.Println("Hello World")

  // Variables are statically typed, which means their type can't change
// Variable names must start with a letter and may contain letters, numbers
// or the _

// An int is a positive or negative number without decimals
// Versions
// uint8 : unsigned  8-bit integers (0 to 255)
// uint16 : unsigned 16-bit integers (0 to 65535)
// uint32 : unsigned 32-bit integers (0 to 4294967295)
// uint64 : unsigned 64-bit integers (0 to 18446744073709551615)
// int8 : signed  8-bit integers (-128 to 127)
// int16 : signed 16-bit integers (-32768 to 32767)
// int32 : signed 32-bit integers (-2147483648 to 2147483647)
// int64 : signed 64-bit integers (-9223372036854775808 to
// 9223372036854775807)

var age int = 27;

var favNum float64 = 1.6180339


fmt.Println(age, favNum)

var numOne = 1.000
var num99 = .999
fmt.Println(numOne - num99)

// Arithmetic Operators : +, -, *, /, %

	fmt.Println("6 + 4 =", 6 + 4)
	fmt.Println("6 - 4 =", 6 - 4)
	fmt.Println("6 * 4 =", 6 * 4)
	fmt.Println("6 / 4 =", 6 / 4)
	fmt.Println("6 % 4 =", 6 % 4)

}
