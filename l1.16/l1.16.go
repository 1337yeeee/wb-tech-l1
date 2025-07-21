package main

import "fmt"

func quickSort(arr []int) []int {
	// Массив из 0 или 1 элемента уже отсортирован
	if len(arr) < 2 {
		return arr
	}

	// Выбираем опорный элемент – первый элемент массива
	pivot := arr[0]

	// Срезы для элементов меньше или равных / больше опорному
	var left, right []int

	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	// Рекурсивно сортируем и объединяем результат
	return append(append(quickSort(left), pivot), quickSort(right)...)
}

func main() {
	arr := []int{10, 5, 2, 3, 1, 1, 9, 1, 4, 6}
	fmt.Println("До сортировки:", arr)
	sorted := quickSort(arr)
	fmt.Println("После сортировки:", sorted)
}
