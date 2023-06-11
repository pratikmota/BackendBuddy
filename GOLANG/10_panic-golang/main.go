package main

import (
	"fmt"
	"os"
)

func main() {

	//The panic built-in function stops normal execution of the current goroutine.
	//When a function F calls panic, normal execution of F stops immediately.
	_, err := os.Create("/tmp/file")
	if err != nil {
		fmt.Println("Error in os.Create")
		panic(err) // Stop Main function
	}

	panic("There is problem, Stop execution of program")
}
