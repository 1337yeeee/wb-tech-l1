// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action
// от родительской структуры Human (аналог наследования).
// Подсказка: используйте композицию (embedded struct),
// чтобы Action имел все методы Human.

package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (human *Human) SayHello() {
	fmt.Printf("Hello! My name is %s. I'm %d\n", human.name, human.age)
}

func (human *Human) CelebrateBirthday() {
	human.age += 1
	fmt.Printf("Today is my birthday! I am %d now!\n", human.age)
}

type Action struct {
	Human // embedded struct
	job   string
}

func (action *Action) DescribeJob() {
	fmt.Printf("My job is %s\n", action.job)
}

var builder = Action{Human{"Tom", 35}, "builder"}

func main() {
	builder.SayHello()
	builder.CelebrateBirthday()
	builder.DescribeJob()
}
