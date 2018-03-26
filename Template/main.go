// Template record basic Go structure
package main

// Import multiple packages using ()
import (
	"fmt"
)

// Constant declaration
// const IDENTIFIER [type] = value
const CHAR byte = 'C'
const LN2 = 0.693147180559945309417232121458176568075500134360255254120680009
const (
	ZERO = iota
	ONE  = iota
	TWO  = iota
	THREE
	FOUR
)

// Variable initialization
// var identifier [type] = value
var v int = 5
var (
	w int = 10
	x     = 17.6
)
var y, z = true, "four"

/*
signed integer:
int, int8, int16, int32, int64

unsigned integer:
uint, uint8, uint16, uint32, uint64

pointer: uintptr

rune -> int32
byte -> uint8

floating-point:
float32, float64

complex number:
complex64, complex128
*/

// Type definition
type T struct{}

// Init function run before main
func init() {

}

func main() {
	var a float64 // Variable declaration and set to default value
	b := 1e6      // Short-hand assigment only available in function
	a, b = b, a   // Swap value
	p := new(T)   // Create variable type T, initialzed to default, return address
	Func1()       // Function call
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T = %[1]v\n", a)
	fmt.Println(p)
	fmt.Println("Program " + "end here")
}

func (t T) Method1() {

}

// Function declaration and definition
// func Identifier(parameter) return [type]
func Func1() {
	fmt.Println("Start...")
}
