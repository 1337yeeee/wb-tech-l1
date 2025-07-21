package main

import "fmt"

// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа x из массива,
// во второй – результат операции x*2.
// После этого данные из второго канала должны выводиться в stdout.
// То есть, организуйте конвейер из двух этапов с горутинами: генерация чисел и их обработка.
// Убедитесь, что чтение из второго канала корректно завершается.

func main() {
	inputChannel := make(chan int)
	outputChannel := make(chan int)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 10, 9, 11, 17, 15, 13, 12}

	// заполнение первого канала числами из массива numbers
	go func() {
		for _, x := range numbers {
			inputChannel <- x
		}
		close(inputChannel)
	}()

	// чтение чисел из канала inputChannel и запись x*2 в канал outputChannel
	go func() {
		for x := range inputChannel {
			outputChannel <- x * 2
		}
		close(outputChannel)
	}()

	// чтение и печать чисел из канала outputChannel
	for result := range outputChannel {
		fmt.Println(result)
	}
}
