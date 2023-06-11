package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	/*
	 1) filepath.Abs
	 2) filepath.Ext(path)
	 3) filepath.Base(path
	 4) filepath.Dir(path)
	*/
	path, err := filepath.Abs("./main.go")
	if err != nil {
		fmt.Printf("error in getting path, err:%v", err)
	}
	fmt.Println(path)
	// Get Extension of file
	ext := filepath.Ext(path)
	fmt.Println("File extension:", ext)           // File extension: .go
	fmt.Printf("Base: %s\n", filepath.Base(path)) // main.go
	// Dir: E:\YOUTUBE\CODE\BackendBuddy\GOLANG\9_get-file-path
	fmt.Printf("Dir: %s\n", filepath.Dir(path))
}
