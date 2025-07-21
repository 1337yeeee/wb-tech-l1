package main

import (
	"fmt"
	"sync"
)

// Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде
// (т.е. из нескольких горутин).
// По завершению программы структура должна выводить итоговое значение счётчика.
// Подсказка: вам понадобится механизм синхронизации,
// например, sync.Mutex или sync/Atomic для безопасного инкремента.

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println("Итоговое значение:", counter.Value())
	// Итоговое значение: 1000

	newCounter := Counter{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 5; j++ {
				newCounter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Println("Итоговое значение:", newCounter.Value())
	// Итоговое значение: 500
}
