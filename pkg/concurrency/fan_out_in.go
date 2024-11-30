package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		time.Sleep(time.Millisecond * 100) // Имитация работы
		results <- job * 2
	}
}

func RunFanOutFanIn() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	var wg sync.WaitGroup

	// Fan-Out: Запускаем несколько горутин
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results)
		}(w)
	}

	// Отправляем задания
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	// Ждем завершения всех воркеров
	wg.Wait()
	close(results)

	// Fan-In: Читаем результаты
	for result := range results {
		fmt.Println("Result:", result)
	}
}
