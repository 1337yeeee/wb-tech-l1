package main

import "fmt"

// Поменять местами два числа без использования временной переменной.
// Подсказка: примените сложение/вычитание или XOR-обмен.

func main() {
	var a int
	var b int

	a = 15
	b = 93
	fmt.Println("Initial values")
	fmt.Printf("a: %d; b: %d\n", a, b)

	a = a + b
	b = a - b
	a = a - b
	fmt.Println("Addition / subtraction")
	fmt.Printf("a: %d; b: %d\n", a, b)

	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Println("XOR")
	fmt.Printf("a: %d; b: %d\n", a, b)

	a, b = b, a
	fmt.Println("Go swap")
	fmt.Printf("a: %d; b: %d\n", a, b)
}
