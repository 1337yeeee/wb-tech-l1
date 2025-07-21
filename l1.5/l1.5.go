// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала – читать эти значения.
// По истечении N секунд программа должна завершаться.
// Подсказка: используйте time.After или таймер для ограничения времени работы.

package main

import (
	"fmt"
	"time"
)

func channelFiller(channel chan int) {
	for i := 0; ; i++ {
		channel <- i
		time.Sleep(time.Duration(200) * time.Millisecond)
	}
}

func channelConsumer(channel chan int) {
	for number := range channel {
		fmt.Printf("Consumed :-> %d\n", number)
	}
}

func main() {
	var N int
	fmt.Print("Enter amount of seconds (N): ")
	fmt.Scan(&N)

	channel := make(chan int)

	go channelFiller(channel)
	go channelConsumer(channel)

	<-time.After(time.Duration(N) * time.Second)
	fmt.Printf("Timer for %d seconds is up\n", N)
	close(channel)
}
