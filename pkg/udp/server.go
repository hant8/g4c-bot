package tcp

import (
	"fmt"
	"log"
	"net"
)

func StartServer(port string) {
	addr, err := net.ResolveUDPAddr("udp", ":"+port)
	if err != nil {
		panic(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}
	defer func(conn *net.UDPConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("Критическая ошибка при закрытии ListenUDP: %v", err)
		}
	}(conn)

	fmt.Printf("UDP сервер слушает на порту %s\n", port)

	buffer := make([]byte, 1024)
	for {
		n, clientAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			continue
		}

		fmt.Printf("Получено от %s: %s\n", clientAddr, string(buffer[:n]))

		_, err = conn.WriteToUDP([]byte("Принято\n"), clientAddr)
		if err != nil {
			fmt.Println("Ошибка отправки:", err)
		}
	}
}
