package main

import (
	"fmt"
)

// Удалить i-ый элемент из слайса. Продемонстрируйте корректное удаление без утечки памяти.
// Подсказка: можно сдвинуть хвост слайса на место удаляемого элемента
// (copy(slice[i:], slice[i+1:])) и уменьшить длину слайса на 1.

func removeElement(slice []int, i int) []int {
	// Проверка корректности индекса
	if i < 0 || i >= len(slice) {
		return slice
	}

	// Сдвигаем элементы после i-го
	copy(slice[i:], slice[i+1:])

	// Обнуляем последний элемент (для предотвращения утечки),
	// например, если слайс содержит указатели на структуры
	slice[len(slice)-1] = 0

	// Возвращаем срез без последнего элемента
	return slice[:len(slice)-1]
}

func main() {
	test()
}

func test() {
	test1 := []int{10, 20, 30, 40, 50}
	result1 := removeElement(test1, 2)
	fmt.Printf("Тест 1 - Удаление из середины:\nИсходный: %v\nРезультат: %v\nОжидаемый: [10 20 40 50]\n\n", []int{10, 20, 30, 40, 50}, result1)

	test2 := []int{10, 20, 30, 40, 50}
	result2 := removeElement(test2, 0)
	fmt.Printf("Тест 2 - Удаление первого элемента:\nИсходный: %v\nРезультат: %v\nОжидаемый: [20 30 40 50]\n\n", []int{10, 20, 30, 40, 50}, result2)

	test3 := []int{10, 20, 30, 40, 50}
	result3 := removeElement(test3, 4)
	fmt.Printf("Тест 3 - Удаление последнего элемента:\nИсходный: %v\nРезультат: %v\nОжидаемый: [10 20 30 40]\n\n", []int{10, 20, 30, 40, 50}, result3)

	test4 := []int{10, 20, 30, 40, 50}
	result4 := removeElement(test4, 10)
	fmt.Printf("Тест 4 - Индекс за границами:\nИсходный: %v\nРезультат: %v\nОжидаемый: [10 20 30 40 50]\n\n", []int{10, 20, 30, 40, 50}, result4)

	test5 := []int{10, 20, 30, 40, 50}
	result5 := removeElement(test5, -1)
	fmt.Printf("Тест 5 - Отрицательный индекс:\nИсходный: %v\nРезультат: %v\nОжидаемый: [10 20 30 40 50]\n\n", []int{10, 20, 30, 40, 50}, result5)

	test6 := []int{}
	result6 := removeElement(test6, 0)
	fmt.Printf("Тест 6 - Пустой слайс:\nИсходный: %v\nРезультат: %v\nОжидаемый: []\n\n", []int{}, result6)

	test7 := []int{100}
	result7 := removeElement(test7, 0)
	fmt.Printf("Тест 7 - Слайс из одного элемента:\nИсходный: %v\nРезультат: %v\nОжидаемый: []\n\n", []int{100}, result7)

	test8 := []int{10, 20, 30, 40, 50}
	originalCap := cap(test8)
	result8 := removeElement(test8, 2)
	fmt.Printf("Тест 8 - Проверка capacity:\nИсходная capacity: %d\nCapacity после удаления: %d\nОжидаемый результат: capacity не изменилась\n\n", originalCap, cap(result8))
}
