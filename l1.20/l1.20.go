package main

import "fmt"

// Разработать программу, которая переворачивает порядок слов в строке.
// Пример: входная строка:
// «snow dog sun», выход: «sun dog snow».
// Считайте, что слова разделяются одиночным пробелом.
// Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».

func ReversWords(s string) string {
	sr := []rune(s)

	reverse(sr, 0, len(sr)-1)

	start := 0
	for i := 0; i <= len(sr); i++ {
		if i == len(sr) || sr[i] == ' ' {
			reverse(sr, start, i-1)
			start = i + 1
		}
	}

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
	var theString, revertedString string

	theString = "главрыба"
	revertedString = ReversWords(theString)
	fmt.Printf("Строка:\t\t%s\nПеревернутая:\t%s\n\n", theString, revertedString)
	//   Строка:		главрыба
	//   Перевернутая:	главрыба

	theString = "snow dog sun"
	revertedString = ReversWords(theString)
	fmt.Printf("Строка:\t\t%s\nПеревернутая:\t%s\n\n", theString, revertedString)
	//   Строка:		snow dog sun
	//   Перевернутая:	sun dog snow

	theString = "Разработать программу, которая переворачивает подаваемую на вход строку."
	revertedString = ReversWords(theString)
	fmt.Printf("Строка:\t\t%s\nПеревернутая:\t%s\n\n", theString, revertedString)
	//   Строка:		Разработать программу, которая переворачивает подаваемую на вход строку.
	//   Перевернутая:	строку. вход на подаваемую переворачивает которая программу, Разработать

	theString = "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	revertedString = ReversWords(theString)
	fmt.Printf("Строка:\t\t%s\nПеревернутая:\t%s\n\n", theString, revertedString)
	//   Строка:		Lorem ipsum dolor sit amet, consectetur adipiscing elit.
	//   Перевернутая:	elit. adipiscing consectetur amet, sit dolor ipsum Lorem
}
