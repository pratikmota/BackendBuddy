package main

import (
	"fmt"
)

func main() {

    funcs := []func() {
        BackendBuddy1,
        BackendBuddy2,
        BackendBuddy3,
    }

    for _, f := range funcs{
        f()
    }

}

func BackendBuddy1() {
    fmt.Println("Backend Buddy 1")
}

func BackendBuddy2() {
    fmt.Println("Backend Buddy 2")
}

func BackendBuddy3() {
    fmt.Println("Backend Buddy 3")
}