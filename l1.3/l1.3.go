// Реализовать постоянную запись данных в канал (в главной горутине).
// Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.
// Программа должна принимать параметром количество воркеров и при старте создавать
// указанное число горутин-воркеров.

// Для корректного завершения горутин-воркеров по SIGINT (Ctrl+C)
// используем context + WaitGroup + канал сигналов.
// context.WithCancel передаёт сигнал отмены всем воркерам через ctx.Done().
// WaitGroup гарантирует, что программа дождётся их остановки.
// Канал с signal.Notify ловит SIGINT, инициируя отмену.
// Этот подход — идиоматичный в Go: он контролируемый и гибкий:
// позволяет управлять таймаутами и отменой (cancel)

package main

import (
	"fmt"
	"math/rand/v2"
)

const MaxOutstanding = 10

func channelFiller(channel chan int) {
	// В бесконечном цикле наполняем канал случайными числами
	for {
		channel <- rand.IntN(100)
	}
}

func channelConsumer(index int, channel chan int) {
	// Из канала получаем чило, выводим его в stdout
	for number := range channel {
		fmt.Printf("Consm{%d} :-> %d\n", index, number)
	}
}

func main() {
	var N int
	fmt.Print("Enter amount of workers (N): ")
	fmt.Scan(&N)

	var the_channel = make(chan int)
	go channelFiller(the_channel)

	for index := range N {
		go channelConsumer(index, the_channel)
	}

	for {
	}
}
