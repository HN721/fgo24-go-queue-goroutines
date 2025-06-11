package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(queue <-chan int, done chan<- bool) {
	for msg := range queue {
		sleepTime := time.Duration(rand.Intn(3)+1) * time.Second
		time.Sleep(sleepTime)
		fmt.Printf("Proses queue %d\n", msg)
		fmt.Println("Waktu selesai ", sleepTime)
	}
	done <- true
}

func main() {

	queue := make(chan int)
	done := make(chan bool)

	count := rand.Intn(5) + 1
	fmt.Printf("Jumlah antrian: %d\n", count)

	go worker(queue, done)

	for x := range count {
		queue <- x + 1
	}

	close(queue)
	<-done
}
