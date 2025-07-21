// Программа должна корректно завершаться по нажатию Ctrl+C (SIGINT).
// Выберите и обоснуйте способ завершения работы всех горутин-воркеров при получении сигнала прерывания.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		// Проверяем завершение контекста
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("Context error: %s\n", err.Error())
			}
			fmt.Printf("Worker {%d} finished\n", id)
			return
		// Выполняем основую функцию
		default:
			fmt.Printf("Working {%d}\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// Создаем воркеров, передаем им контекст и WaitGroup,
	// чтобы отложить завершнеие работы воркера
	numWorkers := 3
	for id := range numWorkers {
		wg.Add(1)
		go worker(ctx, id, &wg)
	}

	// Создаем канал для сигналов
	// Перехватываем SIGINT в sigChan
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Оживаем получение SIGINT
	<-sigChan
	fmt.Printf("Recived SIGING. Interupting...\n")

	// Завершаем контекст, дожидаемся завершения горутин
	cancel()
	wg.Wait()

	fmt.Printf("Shutted down\n")
}
