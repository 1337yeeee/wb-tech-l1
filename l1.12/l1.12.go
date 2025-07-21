package main

import "fmt"

// Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree").
// Создать для неё собственное множество.
// Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

func UniqueWords(words []string) []string {
	// множество значений для проверки на уникальность
	set := make(map[string]struct{})
	// результирующий слайс
	result := []string{}

	for _, word := range words {
		// проверка: встречалось ли слово раньше
		if _, exists := set[word]; !exists {
			set[word] = struct{}{}
			result = append(result, word)
		}
	}

	return result
}

func CallUniqueWordsAndPrintResult(words []string) {
	result := UniqueWords(words)
	fmt.Printf("Входные слова: %v\n", words)
	fmt.Printf("Результат: %v\n\n", result)
}

func main() {
	var words []string

	words = []string{"cat", "cat", "dog", "cat", "tree"}
	CallUniqueWordsAndPrintResult(words)
	// Входные слова: [cat cat dog cat tree]
	// Результат: [cat dog tree]

	words = []string{"cat", "cat", "cat"}
	CallUniqueWordsAndPrintResult(words)
	// Входные слова: [cat cat cat]
	// Результат: [cat]

	words = []string{"cat", "dog", "spam", "egg", "foo", "bar"}
	CallUniqueWordsAndPrintResult(words)
	// Входные слова: [cat dog spam egg foo bar]
	// Результат: [cat dog spam egg foo bar]

	words = []string{}
	CallUniqueWordsAndPrintResult(words)
	// Входные слова: []
	// Результат: []
}
