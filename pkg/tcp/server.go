package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

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
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка принятия соединения:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Критическая ошибка при закрытии Connection: %v", err)
		}
	}(conn)

	fmt.Println("Новое соединение:", conn.RemoteAddr())
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
}
