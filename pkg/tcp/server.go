package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const workerCount int = 4

func StartServer(port string) {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}

	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatalf("Критическая ошибка при закрытии listener: %v", err)
		}
	}(listener)

	fmt.Printf("Сервер слушает на порту %s\n", port)
	jobs := initializeWorkerPool()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка принятия соединения:", err)
			continue
		}

		jobs <- conn
	}
}

func initializeWorkerPool() chan<- net.Conn {
	jobs := make(chan net.Conn, 100)

	for i := 0; i < workerCount; i++ {
		go worker(i, jobs)
	}

	return jobs
}

func worker(id int, jobs <-chan net.Conn) {
	for conn := range jobs {
		fmt.Printf("Воркер %d обрабатывает соединение от %s\n", id, conn.RemoteAddr())
		_, err := conn.Write([]byte("Привет, клиент!\n"))
		if err != nil {
			return
		}

		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println("Получено:", scanner.Text())
			_, err := conn.Write([]byte("Принято: " + scanner.Text() + "\n"))
			if err != nil {
				fmt.Println("Ошибка записи по соединеннию:", conn.RemoteAddr())
				return
			}
		}

		fmt.Println("Закрытие соединения", conn.RemoteAddr())

		err = conn.Close()
		if err != nil {
			return
		}
	}
}
