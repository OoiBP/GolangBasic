// Loop record repetitive task
// Go has no while loop
package main

import "fmt"

func main() {
	// for [initialization]; [condition]; [post] {}
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// for [condition] {}
	var x = 3
	for x < 6 {
		fmt.Println(x)
		x++
	}

	// for {} infinity loop
	for {
	}

	// for index, argument := range slice {}
	// use blank identifier _ for unwanted variable
}
