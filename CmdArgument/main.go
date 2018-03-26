// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

// Build program
// execute program with ./main.exe Arguments
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	fmt.Println("stop")
}
