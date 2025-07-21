package main

import "fmt"

// Разработать программу, которая переворачивает подаваемую на вход строку.
// Например: при вводе строки «главрыба» вывод должен быть «абырвалг».
// Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
// то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).

func RevertString(s string) string {
	var reveted []rune
	sr := []rune(s)

	for i := len(sr) - 1; i >= 0; i-- {
		reveted = append(reveted, sr[i])
	}

	return string(reveted)
}

func main() {
	var theString, revertedString string

	theString = "главрыба"
	revertedString = RevertString(theString)
	fmt.Printf("Строка:\t\t%s\nПеревернутая:\t%s\n\n", theString, revertedString)
	//   Строка:		главрыба
	//   Перевернутая:	абырвалг

	theString = "Разработать программу, которая переворачивает подаваемую на вход строку."
	revertedString = RevertString(theString)
	fmt.Printf("Строка:\t\t%s\nПеревернутая:\t%s\n\n", theString, revertedString)
	//   Строка:		Разработать программу, которая переворачивает подаваемую на вход строку.
	//   Перевернутая:	.укортс дохв ан юумеавадоп теавичаровереп яароток ,уммаргорп ьтатобарзаР

	theString = "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	revertedString = RevertString(theString)
	fmt.Printf("Строка:\t\t%s\nПеревернутая:\t%s\n\n", theString, revertedString)
	//   Строка:		Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	//   Перевернутая:	.tile gnicsipida rutetcesnoc ,tema tis rolod muspi meroL
}
