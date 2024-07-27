package main

import "fmt"

// Varibales in Go

// bool: a boolean value, either true or false
// string: a sequence of characters
// int: a signed integer
// float64: a floating-point number
// byte: exactly what it sounds like: 8 bits of data

func main() {
	// Normal Way of Variable Declaration
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
	// %v  - default format (for smsSendingLimit, hasPermission)
	// %.2f - float with 2 decimal places (for costPerSMS)
	// %q  - quoted string (for username)
	// \n  - newline

	// Shorthand way to declare variables
	empty := ""
	// Same as var empty string
	// := can't be used outside any function i.e. global scope
}
