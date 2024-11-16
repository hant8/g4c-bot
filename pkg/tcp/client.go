package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func StartClient(port string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		panic(err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Критическая ошибка при закрытии listener: %v", err)
		}
	}(conn)

	fmt.Println("Соединение установлено. Введите текст:")

	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Ошибка отправки:", err)
			break
		}

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Ответ сервера:", message)
	}
}
