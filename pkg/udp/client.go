package udp

import (
	"fmt"
	"log"
	"net"
)

func StartClient(port string) {
	serverAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		panic(err)
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Критическая ошибка при закрытии DialUDP: %v", err)
		}
	}(conn)

	fmt.Println("Введите текст для отправки:")

	for {
		var text string
		_, err = fmt.Scanln(&text)
		if err != nil {
			return
		}

		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println("Ошибка отправки:", err)
			break
		}

		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Ошибка получения:", err)
			break
		}
		fmt.Println("Ответ от сервера:", string(buffer[:n]))
	}
}
