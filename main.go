package main

import "fmt"

// Функция, которая складывает два числа
func Add(a, b int) int {
    return a + b
}

func main() {
    result := Add(2, 3)
    fmt.Printf("2 + 3 = %d\n", result)
}
