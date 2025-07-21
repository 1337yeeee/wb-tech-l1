package main

// Дана переменная типа int64. Разработать программу,
// которая устанавливает i-й бит этого числа в 1 или 0.
// Пример: для числа 5 (0101₂) установка 1-го бита в 0 даст 4 (0100₂).
// Подсказка: используйте битовые операции (|, &^).

import (
	"errors"
	"fmt"
	"strconv"
)

func SetBit(num int64, i int, bitValue uint) (int64, error) {
	if i <= 0 || i > 64 {
		return 0, errors.New("Недопустимое значение разряда")
	}

	i--

	if bitValue == 1 {
		return (num | (1 << i)), nil
	} else {
		return (num &^ (1 << i)), nil
	}
}

func SetBitAndPrintResult(num int64, i int, bitValue uint) {
	result, err := SetBit(num, i, bitValue)

	fmt.Printf("Число %d; разряд %d; новое значение бита %d; результат: %d\n", num, i, bitValue, result)

	if err != nil {
		fmt.Printf("Ошибка: %v\n\n", err.Error())
		return
	}

	fmt.Printf("Число %s;\nрезультат: %s\n\n", strconv.FormatInt(num, 2), strconv.FormatInt(result, 2))
}

func main() {
	var num int64
	var i int
	var bitValue uint

	// 1 (0001), сбрасываем 1-й бит (самый младший) -> получится 0 (0000)
	num = 1
	i = 1
	bitValue = 0
	SetBitAndPrintResult(num, i, bitValue)

	// 1 (0001), устанавливаем 1-й бит в 1 -> уже установлен, результат равен исходному числу: 1 (0001)
	num = 1
	i = 1
	bitValue = 1
	SetBitAndPrintResult(num, i, bitValue)

	// 5 (0101), сбрасываем 1-й бит -> получим 4 (0100)
	num = 5
	i = 1
	bitValue = 0
	SetBitAndPrintResult(num, i, bitValue)

	// 4 (0100), устанавливаем 1-й бит -> получим 5 (0101)
	num = 4
	i = 1
	bitValue = 1
	SetBitAndPrintResult(num, i, bitValue)

	// 7 (0111), сбрасываем 3-й бит (старший среди 3-х) -> получим 3 (0011)
	num = 7
	i = 3
	bitValue = 0
	SetBitAndPrintResult(num, i, bitValue)

	// 0 (00...01), устанавливаем 64-й бит (бит знака для int64) -> получим отрицательное число: math.MinInt64 = -9223372036854775808
	num = 0
	i = 64
	bitValue = 1
	SetBitAndPrintResult(num, i, bitValue)

	// 6 (0110), устанавливаем 2-й бит -> уже установлен, результат равен исходному числу: 6 (0110)
	num = 6
	i = 2
	bitValue = 1
	SetBitAndPrintResult(num, i, bitValue)

	// 8 (1000), сбрасываем 2-й бит (значение ноль) -> результат равен исходному числу: 8 (1000)
	num = 8
	i = 2
	bitValue = 0
	SetBitAndPrintResult(num, i, bitValue)

	// Ошибка: i == 0, что вне допустимого диапазона [1, 64]
	num = 5
	i = 0
	bitValue = 0
	SetBitAndPrintResult(num, i, bitValue)

	// Ошибка: i == 65, что превышает максимальный разряд 64
	num = 5
	i = 65
	bitValue = 0
	SetBitAndPrintResult(num, i, bitValue)
}
