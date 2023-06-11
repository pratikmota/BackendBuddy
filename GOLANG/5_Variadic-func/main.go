package main

import "fmt"

func Add(values ...int) {
    fmt.Print(values, " ")
    total := 0
    for _, num := range values {
        total += num
    }
    fmt.Println(total)
}
func main() {
    Add(1, 2) // 1+2 = 3
    Add(1, 2, 3) // 1+2+3 = 6

    // Variadic functions for slices
    values := []int{1, 2, 3, 4}
    Add(values...)
}
