package withgroup

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(number int, queue *sync.WaitGroup) {
	defer queue.Done()

	sleepTime := time.Duration(rand.Intn(3)+1) * time.Second
	time.Sleep(sleepTime)
	fmt.Printf("Proses queue %d\n", number)
	fmt.Println("Waktu selesai:", sleepTime)
}

func Withgroup() {

	var queue sync.WaitGroup

	count := rand.Intn(5) + 1
	fmt.Printf("Jumlah antrian: %d\n", count)

	for i := range count {
		queue.Add(1)
		go worker(i, &queue)
	}

	queue.Wait()
}
