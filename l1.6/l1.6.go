package main

// Реализовать все возможные способы остановки выполнения горутины.
// Классические подходы: выход по условию, через канал уведомления,
// через контекст, прекращение работы runtime.Goexit() и др.
// Продемонстрируйте каждый способ в отдельном фрагменте кода.

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

const (
	ColorReset   = "\033[0m"
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorMagenta = "\033[35m"
	ColorCyan    = "\033[36m"
)

// Выход по условию
func workerCondition(stopFlag *bool) {
	for !*stopFlag {
		fmt.Printf("workerCondition %sработает%s\n", ColorGreen, ColorReset)
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Printf("%smworkerCondition%s %sостановлен%s по флагу\n", ColorCyan, ColorReset, ColorMagenta, ColorReset)
}

// Выход через канал уведомлений
func workerChanel(stopChan chan int) {
	for {
		select {
		case <-stopChan:
			fmt.Printf("%smworkerChanel%s %sостановлен%s через канал уведомления\n", ColorCyan, ColorReset, ColorMagenta, ColorReset)
			return
		default:
			fmt.Printf("workerChanel %sработает%s\n", ColorGreen, ColorReset)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// Выход через контекст
func workerContext(ctx context.Context, message string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%smworkerContext%s %sостановлен%s по контексту: %s\n", ColorCyan, ColorReset, ColorMagenta, ColorReset, message)
			return
		default:
			fmt.Printf("workerContext[%s] %sработает%s\n", message, ColorGreen, ColorReset)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// Выход через Goexit
func workerGoexit() {
	defer fmt.Printf("%smworkerGoexit%s %sостановлен%s через Goexit\n", ColorCyan, ColorReset, ColorMagenta, ColorReset)

	fmt.Printf("workerGoexit %sработает%s\n", ColorGreen, ColorReset)

	runtime.Goexit()
}

// Выход через Goexit с предварительной паникой
func workerPanicGoexit() {
	defer fmt.Printf("%smworkerPanicGoexit%s %sостановлен%s через Goexit (Без паники!)\n", ColorCyan, ColorReset, ColorMagenta, ColorReset)
	defer runtime.Goexit()

	fmt.Printf("workerPanicGoexit %sработает%s\n", ColorGreen, ColorReset)

	panic("A-A-A-A-A-A-A-A-A-A-A-A-A-A-A")
}

// Выход через панику и ее перехват
func workerPanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("workerPanic: перехвачена паника: %s%v%s\n", ColorRed, r, ColorReset)
		}
		fmt.Printf("%smworkerPanic%s %sостановлен%s после перехвата паники\n", ColorCyan, ColorReset, ColorMagenta, ColorReset)
	}()

	fmt.Printf("workerPanic %sработает%s\n", ColorGreen, ColorReset)

	panic("A-A-A-A-A-A-A-A-A-A-A-A-A-A-A")

	fmt.Println("Это не будет напечатано") // unreachable code
}

func main() {
	var wg sync.WaitGroup
	stopFlag := false
	stopChan := make(chan int)
	ctxCancel, cancel := context.WithCancel(context.Background())
	ctxCancelTimeout, cancelTimeout := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()
	defer cancelTimeout()

	wg.Add(1)
	go func() {
		defer wg.Done()
		workerCondition(&stopFlag)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		workerChanel(stopChan)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		workerContext(ctxCancel, "cancel")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		workerContext(ctxCancelTimeout, "timeout")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		workerGoexit()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		workerPanicGoexit()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		workerPanic()
	}()

	fmt.Print("created\n")

	time.Sleep(1 * time.Second)

	stopFlag = true
	stopChan <- 1
	cancel()

	wg.Wait()
	fmt.Print("Wait group done!\n")
}
