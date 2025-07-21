package main

import (
	"fmt"
	"sync"
)

// Реализовать безопасную для конкуренции запись данных в структуру map.
// Подсказка: необходимость использования синхронизации
// (например, sync.Mutex или встроенная concurrent-map).
// Проверьте работу кода на гонки (util go run -race).

type SafeMap struct {
	mu   sync.Mutex
	data map[string]int
}

func (sm *SafeMap) Set(key string, val int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = val
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	val, ok := sm.data[key]
	return val, ok
}

func main() {
	safeMap := SafeMap{data: make(map[string]int)}
	var wg sync.WaitGroup

	// Запись в SafeMap
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n%10)
			safeMap.Set(key, n)
		}(i)
	}

	wg.Wait()

	// Выод результата
	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key%d", i)
		val, ok := safeMap.Get(key)
		if ok {
			fmt.Printf("%s: %d\n", key, val)
		}
	}
}
