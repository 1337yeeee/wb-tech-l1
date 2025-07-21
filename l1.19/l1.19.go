package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на вход строку.
// Например: при вводе строки «главрыба» вывод должен быть «абырвалг».
// Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
// то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).

func Reverse(s string) string {
	sr := []rune(s)

	reverse(sr, 0, len(sr)-1)

	return string(sr)
}

func reverse(r []rune, start, end int) {
	for start < end {
		r[start], r[end] = r[end], r[start]
		start++
		end--
	}
}

func main() {
	var theString, reversedString string

	theString = "главрыба"
	reversedString = Reverse(theString)
	fmt.Printf("Строка:\t\t%s\nПеревернутая:\t%s\n\n", theString, reversedString)
	//   Строка:		главрыба
	//   Перевернутая:	абырвалг

	theString = "Разработать программу, которая переворачивает подаваемую на вход строку."
	reversedString = Reverse(theString)
	fmt.Printf("Строка:\t\t%s\nПеревернутая:\t%s\n\n", theString, reversedString)
	//   Строка:		Разработать программу, которая переворачивает подаваемую на вход строку.
	//   Перевернутая:	.укортс дохв ан юумеавадоп теавичаровереп яароток ,уммаргорп ьтатобарзаР

	theString = "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	reversedString = Reverse(theString)
	fmt.Printf("Строка:\t\t%s\nПеревернутая:\t%s\n\n", theString, reversedString)
	//   Строка:		Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	//   Перевернутая:	.tile gnicsipida rutetcesnoc ,tema tis rolod muspi meroL
}
