package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "0.0.0.0:1080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}

		fmt.Printf("连接成功! clent:%v \n", client.RemoteAddr())
		go process(client)
	}
}

func process(conn net.Conn) {
	defer func() {
		conn.Close()
		fmt.Printf("连接断开! clent:%v \n", conn.RemoteAddr())
	}()

	reader := bufio.NewReader(conn)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		_, err = conn.Write([]byte{b})
		if err != nil {
			break
		}
	}
}
