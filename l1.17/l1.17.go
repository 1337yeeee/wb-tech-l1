package main

import (
	"fmt"
	"math/rand/v2"
)

// Реализовать алгоритм бинарного поиска встроенными методами языка.
// Функция должна принимать отсортированный слайс и искомый элемент,
// возвращать индекс элемента или -1, если элемент не найден.
// Подсказка: можно реализовать рекурсивно или итеративно, используя цикл for.

const RANDOM_EDGE = 100000

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]

	var left, right []int

	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func numbersGenerator(length int) []int {
	var numbers []int

	for i := 1; i < length; i++ {
		number := rand.IntN(RANDOM_EDGE)
		numbers = append(numbers, number)
	}

	return quickSort(numbers)
}

func binarySearch(numbers []int, target int) int {
	return binSearch(numbers, target, 0, len(numbers)-1)
}

func binSearch(numbers []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	if numbers[mid] == target {
		return mid
	}
	if numbers[mid] > target {
		return binSearch(numbers, target, left, mid-1)
	}
	return binSearch(numbers, target, mid+1, right)
}

func main() {
	var numbers []int
	var index int
	var target int

	numbers = []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	}
	target = 11
	index = binarySearch(numbers, target)
	fmt.Println("Числа:", numbers)
	fmt.Println("Цель:", target, "Индекс:", index, "\n\n")
	// Числа: [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]
	// Цель: 11 Индекс: 10

	target = 2
	index = binarySearch(numbers, target)
	fmt.Println("Числа:", numbers)
	fmt.Println("Цель:", target, "Индекс:", index, "\n\n")
	// Числа: [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]
	// Цель: 2 Индекс: 1

	target = 1
	index = binarySearch(numbers, target)
	fmt.Println("Числа:", numbers)
	fmt.Println("Цель:", target, "Индекс:", index, "\n\n")
	// Числа: [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]
	// Цель: 1 Индекс: 0

	target = 20
	index = binarySearch(numbers, target)
	fmt.Println("Числа:", numbers)
	fmt.Println("Цель:", target, "Индекс:", index, "\n\n")
	// Числа: [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]
	// Цель: 20 Индекс: 19

	target = 21
	index = binarySearch(numbers, target)
	fmt.Println("Числа:", numbers)
	fmt.Println("Цель:", target, "Индекс:", index, "\n\n")
	// Числа: [1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20]
	// Цель: 21 Индекс: -1

	numbers = numbersGenerator(rand.IntN(100) + 1)
	target = numbers[20]
	index = binarySearch(numbers, target)
	fmt.Println("Числа:", numbers)
	fmt.Println("Цель:", target, "Индекс:", index, "\n\n")
	// Числа: [<отсортированный массив длиной 1-100 из целых чисел (значения ≥ 0)>]
	// Цель: <число под индексом 20> Индекс: 20 (или меньше, так как могут быть одинаковые значения)

	target = -100
	index = binarySearch(numbers, target)
	fmt.Println("Числа:", numbers)
	fmt.Println("Цель:", target, "Индекс:", index, "\n\n")
	// Числа: [<отсортированный массив длиной 1-100 из целых чисел (значения ≥ 0)>]
	// Цель: -100 Индекс: -1 (числа в массиве – неотрицательные)
}
