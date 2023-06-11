package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	num1:= 10
	// 10 int
	fmt.Println(num1, reflect.TypeOf(num1)) 
	// CONVERT INT to STRING
	str:= strconv.Itoa(num1)  
	// 10 string
	fmt.Println(str, reflect.TypeOf(str)) 

	// STRING to INT
	value, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	// 10 int
	fmt.Println(value, reflect.TypeOf(value)) 
}