package main

import "fmt"

func main() {
	// Type Sizes in Go

	// Integers (whole numbers, no decimal)
	var i int = 42
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807

	// Unsigned Integers (positive whole numbers, no decimal)
	var ui uint = 42
	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615

	// Floating-point numbers (signed decimal numbers)
	var f32 float32 = 3.14159265358979323846
	var f64 float64 = 3.14159265358979323846

	// Complex numbers (rarely used)
	var c64 complex64 = 1 + 2i
	var c128 complex128 = 1 + 2i

	// Printing examples
	fmt.Printf("int: %v, max int8: %v, max int16: %v, max int32: %v, max int64: %v\n", i, i8, i16, i32, i64)
	fmt.Printf("uint: %v, max uint8: %v, max uint16: %v, max uint32: %v, max uint64: %v\n", ui, ui8, ui16, ui32, ui64)
	fmt.Printf("float32: %v, float64: %v\n", f32, f64)
	fmt.Printf("complex64: %v, complex128: %v\n", c64, c128)

	// Converting between types
	temperatureFloat := 88.26
	temperatureInt := int64(temperatureFloat)
	fmt.Printf("\nFloat: %v, Converted to int64: %v\n", temperatureFloat, temperatureInt)

// Notes:
// 1. The "default" int and uint types are 32 or 64-bit depending on the environment.
// 2. Standard sizes to use (unless you have specific performance needs):
//    - int
//    - uint
//    - float64
//    - complex128
// 3. Casting a float to an integer truncates the decimal portion.
// 4. Be cautious when converting between types to avoid overflow or loss of precision.
