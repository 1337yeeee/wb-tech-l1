package main

import (
	"fmt"
	"reflect"
)

func TypeChecker(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Println("Целое число, int")
	case string:
		fmt.Println("Строка, string")
	case bool:
		fmt.Println("Булево значение, bool")
	case chan any:
		fmt.Println("Канал, chan any")
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			fmt.Printf("Канал, chan %v\n", reflect.TypeOf(v).Elem())
			return
		}
		fmt.Printf("Тип не поддерживается: %T\n", val)
	}
}

func main() {

	i := 1
	s := "qwe"
	b := true
	c := make(chan interface{})
	ch := make(chan float64)
	f := 3.14
	r := 'r'
	img := complex(3, 2)

	TypeChecker(i)              // Целое число, int
	TypeChecker(s)              // Строка, string
	TypeChecker(b)              // Булево значение, bool
	TypeChecker(c)              // Канал, chan any
	TypeChecker(ch)             // Канал, chan float64
	TypeChecker(f)              // Тип не поддерживается: float64
	TypeChecker(r)              // Тип не поддерживается: int32
	TypeChecker(img)            // Тип не поддерживается: complex128
	TypeChecker(9)              // Целое число, int
	TypeChecker([]int{1, 2, 3}) // Тип не поддерживается: []int
}
