package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id   int
	Name string
}

func main() {
	var jsonInput = []byte("[{\"Id\":1,\"Name\":\"Backend\"},{\"Id\":2,\"Name\":\"Buddy\"}]")
	var students []Student
	//func json.Unmarshal(data []byte, v any) error
	//Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
	err := json.Unmarshal(jsonInput, &students)
	if err != nil {
		fmt.Println(err)
		return
	}
	// [{1 Backend} {2 Buddy}]
	fmt.Println(students)

	for _, student := range students {
		fmt.Printf("ID:%d, Name:%s \n", student.Id, student.Name)
	}
	/*
		ID:1, Name:Backend
		ID:2, Name:Buddy
	*/
}
