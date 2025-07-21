// Написать программу, которая конкурентно рассчитает значения квадратов чисел,
// взятых из массива [2,4,6,8,10], и выведет результаты в stdout.
// Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.

package main

import (
	"fmt"
	"sync"
)

func main() {
	// Задаем массив, обявляем WaitGroup
	var numbers = []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup

	for _, num := range numbers {
		// для каждого воркера увеличиваем WaitGroup
		wg.Add(1)

		go func() {
			// Откладываем уменьшенее WaitGroup
			defer wg.Done()
			square := num * num
			fmt.Println(square)
		}()
	}

	// Дожидаемся завершения всех горутин в WaitGroup
	wg.Wait()
}
