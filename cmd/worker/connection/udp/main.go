package main

import (
	"flag"
	"fmt"
	"go-example/pkg/udp"
)

func main() {
	mode := flag.String("mode", "server", "Запуск в режиме 'server' или 'client'")
	port := flag.String("port", "8080", "Порт для сервера или клиента")
	flag.Parse()

	switch *mode {
	case "server":
		fmt.Println("Запуск UDP-сервера...")
		udp.StartServer(*port)
	case "client":
		fmt.Println("Запуск UDP-клиента...")
		udp.StartClient(*port)
	default:
		fmt.Println("Неизвестный режим. Используйте 'server' или 'client'.")
	}
}
