package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// #1 Print text with \n
	fmt.Println("1. Hello Backend Buddy")

	// #2 Print text with formatting value
	fmt.Printf("%d. Hello Backend Buddy with formating\n",2)

	// #3 Print text and return number of written bytes and error
	byteWritten, err := fmt.Fprint(os.Stdout,"3. Hello Buddy\n")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(byteWritten, " bytes written.\n")

	// #4 Print formated text and return number of written bytes and error
	byteWritten, err = fmt.Fprintf(os.Stdout,"%d. Hello Buddy\n",4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(byteWritten, " bytes written.\n")


	// #5 Scan variable
	var name1 string
	fmt.Scan(&name1)
	fmt.Printf("Name1: %s\n", name1)

	// #6 Scan variable with formatting
	var name2 string
	fmt.Scanf("%s", &name2)
	fmt.Printf("Name2: %s\n", name2)

	// #7 Scan text with datatype with Space
	var st string
	var in int
	r := strings.NewReader("Backend 1 Buddy")
	n, err := fmt.Fscan(r, &st, &in)
	//n, err := fmt.Fscanf(r, "%s %d", &st, &in)

	if err != nil {
        fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
    }
	fmt.Printf("st: %s, in: %d, Fscan word: %d",st,in, n)
} 
	