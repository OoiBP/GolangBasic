package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Read user input
	input := bufio.NewScanner(os.Stdin)

	// if [initialization]; condition {}
	if input.Scan() {
		fmt.Println(input.Text())
	} else {
		fmt.Println("nothing")
	}

	/*
		switch case {
			case condition:
			case condition:
			default:
		}
		no break needed, no fall through
	*/

	/*
		tagless switch
		switch {
			case x > 0:
			default:
			case x < 0:
		}
	*/
}
