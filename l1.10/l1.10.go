package main

import (
	"errors"
	"fmt"
	"math"
)

// Дана последовательность температурных колебаний:
// -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить эти значения в группы с шагом 10 градусов.
// Пример: -20:{-25.4, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20:{24.5}, 30:{32.5}.
// Пояснение:
// 		диапазон -20 включает значения от -20 до -29.9,
// 		диапазон 10 – от 10 до 19.9, и т.д.
// Порядок в подмножествах не важен.

// Из примера к задаче можно заметить, что:
// 		отрицательные диапазоны уходят в меньшую сторону (диапазон -20 включает значения от -20 до -29.9)
//		положительные – в большую (диапазон 10 – от 10 до 19.9)
// Тогда диапазон 0 должен сочитать в себе и положительную природу диапазонов и отприцательную
// Чтобы покрыть все температурную шкалу диапазонами
// Пояснение:
//		диапазон 0 влючает в себя значения от -9.9 до 9.9
// В общем случае диапазон 0 включает в себя значения от (-step + 0.1) до (step - 0.1)

func MakeTemperatureGroups(temeratures []float64, step uint) (map[int][]float64, error) {
	if step == 0 {
		// Для избежания логической ошибки деления на 0
		return nil, errors.New("Шаг диапазона \"0\" не поддерживается")
	}
	groups := make(map[int][]float64)
	var key int

	for _, temp := range temeratures {
		if temp >= 0 {
			key = int(math.Floor(temp/float64(step))) * int(step)
		} else {
			key = int(math.Ceil(temp/float64(step))) * int(step)
		}

		groups[key] = append(groups[key], temp)
	}

	return groups, nil
}

func CallFuncAndPrintResult(temeratures []float64, step uint) {
	fmt.Printf("Шаг: %d\n", step)

	result, err := MakeTemperatureGroups(temeratures, step)
	if err != nil {
		fmt.Printf("Ошибка: %s\n\n", err.Error())
		return
	}

	fmt.Printf("%v\n\n", result)
}

func main() {
	temeratures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -5.5, -6, 3, -19.9, 0, 5.4, 0.5, -0.5}
	var step uint

	step = 10
	CallFuncAndPrintResult(temeratures, step)
	// Шаг: 10
	// map[
	// 		-20:[-25.4 -27 -21]
	// 		-10:[-19.9]
	// 		0:[-5.5 -6 3 0 5.4 0.5 -0.5]
	// 		10:[13 19 15.5]
	// 		20:[24.5]
	// 		30:[32.5]
	// ]

	step = 5
	CallFuncAndPrintResult(temeratures, step)
	// Шаг: 5
	// map[
	// 		-25:[-25.4 -27]
	// 		-20:[-21]
	// 		-15:[-19.9]
	// 		-5:[-5.5 -6]
	// 		0:[3 0 0.5 -0.5]
	// 		5:[5.4]
	// 		10:[13]
	// 		15:[19 15.5]
	// 		20:[24.5]
	// 		30:[32.5]
	// ]

	step = 1
	CallFuncAndPrintResult(temeratures, step)
	// Шаг: 1
	// map[
	// 		-27:[-27]
	// 		-25:[-25.4]
	// 		-21:[-21]
	// 		-19:[-19.9]
	// 		-6:[-6]
	// 		-5:[-5.5]
	// 		0:[0 0.5 -0.5]
	// 		3:[3]
	// 		5:[5.4]
	// 		13:[13]
	// 		15:[15.5]
	// 		19:[19]
	// 		24:[24.5]
	// 		32:[32.5]
	// ]

	step = 0
	CallFuncAndPrintResult(temeratures, step)
	// Шаг: 0
	// Ошибка: Шаг диапазона "0" не поддерживается
}
