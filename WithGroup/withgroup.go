package withgroup

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(queue <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range queue {
		sleepTime := time.Duration(rand.Intn(3)+1) * time.Second
		time.Sleep(sleepTime)
		fmt.Printf("Proses queue %d\n", msg)
		fmt.Println("Waktu selesai", sleepTime)
	}
}

func WithGroup() {

	queue := make(chan int)
	var wg sync.WaitGroup

	count := rand.Intn(5) + 1
	fmt.Printf("Jumlah antrian: %d\n", count)

	wg.Add(1)
	go worker(queue, &wg)

	for x := range count {
		queue <- x + 1
	}

	close(queue)
	wg.Wait()
}
