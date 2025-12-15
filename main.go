package main

import (
	"fmt"
	"runtime"
)

// Функция, которая складывает два числа
func Add(a, b int) int {
	return a + b
}

// Функция, которая возвращает информацию о платформе
func GetPlatformInfo() string {
	return fmt.Sprintf("OS: %s, Architecture: %s", runtime.GOOS, runtime.GOARCH)
}

func main() {
	result := Add(2, 3)
	fmt.Printf("2 + 3 = %d\n", result)
	fmt.Println("Running on:", GetPlatformInfo())
}
