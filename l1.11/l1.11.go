package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов)
// — т.е. вывести элементы, присутствующие и в первом, и во втором.

// Пример:
// A = {1,2,3}
// B = {2,3,4}
// Пересечение = {2,3}

func Intersection(set1 []int, set2 []int) []int {
	set1Map := make(map[int]struct{})
	result := []int{}

	// храним значения из первого множества set1
	for _, x := range set1 {
		set1Map[x] = struct{}{}
	}

	// итерируем второе множество set2
	for _, x := range set2 {
		// проверяем начилие числа в первом множестве set1
		if _, exists := set1Map[x]; exists {
			// если значение из второго множества есть в первом, то добавлем его в result
			result = append(result, x)
		}
	}

	return result
}

// Если в переданных значениях есть дубликаты,
// а возвращаемое значение должно быть Множеством (без дубликатов), то
// добавляем еще одну карту, которая будет хранить добавленные в result значения,
// чтобы не допустить дубликаты
func IntersectionWithoutDuplicates(set1 []int, set2 []int) []int {
	set1Map := make(map[int]struct{})
	result := []int{}
	addedSet := make(map[int]struct{})

	for _, x := range set1 {
		set1Map[x] = struct{}{}
	}

	for _, x := range set2 {
		if _, exists := set1Map[x]; exists {
			if _, added := addedSet[x]; !added {
				result = append(result, x)
				addedSet[x] = struct{}{}
			}
		}
	}

	return result
}

func CallIntersectionAndPrintResult(set1 []int, set2 []int) {
	result := Intersection(set1, set2)
	fmt.Printf("Set  1: %v\n", set1)
	fmt.Printf("Set  2: %v\n", set2)
	fmt.Printf("Result: %v\n\n", result)
}

func main() {
	var set1, set2 []int

	set1 = []int{1, 2, 3}
	set2 = []int{2, 3, 4}
	CallIntersectionAndPrintResult(set1, set2)
	// Set  1: [1 2 3]
	// Set  2: [2 3 4]
	// Result: [2 3]

	set1 = []int{}
	set2 = []int{}
	CallIntersectionAndPrintResult(set1, set2)
	// Set  1: []
	// Set  2: []
	// Result: []

	set1 = []int{1, 2, 3, 4, 5, 6}
	set2 = []int{7, 8, 9, 10}
	CallIntersectionAndPrintResult(set1, set2)
	// Set  1: [1 2 3 4 5 6]
	// Set  2: [7 8 9 10]
	// Result: []

	set1 = []int{5, 6, 2}
	set2 = []int{9, 2}
	CallIntersectionAndPrintResult(set1, set2)
	// Set  1: [5 6 2]
	// Set  2: [9 2]
	// Result: [2]

	set1 = []int{0}
	set2 = []int{0}
	CallIntersectionAndPrintResult(set1, set2)
	// Set  1: [0]
	// Set  2: [0]
	// Result: [0]

	set1 = []int{100}
	set2 = []int{}
	CallIntersectionAndPrintResult(set1, set2)
	// Set  1: [100]
	// Set  2: []
	// Result: []

	set1 = []int{10, 5, 16}
	set2 = []int{16, 10, 5}
	CallIntersectionAndPrintResult(set1, set2)
	// Set  1: [10 5 16]
	// Set  2: [16 10 5]
	// Result: [16 10 5]
}
