package main

//auth 阶段
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)
const(
	socks5Ver = 0x05
	cmdBind = 0x01
	atypIPV4 = 0x01
	atypeHOST = 0x03
	atypeIPV6 = 0x04
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
	err := auth(reader,conn)
	if err!=nil{
		log.Printf("client %v auth failed:%v",conn.RemoteAddr(),err)
	}
	log.Println("auth success")
}

func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	ver,err := reader.ReadByte()
	if err != nil{
		return fmt.Errorf("read ver failed:%w",err)
	}
	if ver != socks5Ver{
		return fmt.Errorf("not supported ver:%v",ver)
	}
	methodSize,err := reader.ReadByte()
	if err!=nil{
		return fmt.Errorf("read methodSize failed:%w",err)
	}
	method := make([]byte,methodSize)
	_,err = io.ReadFull(reader,method)
	if err!=nil{
		return fmt.Errorf("read method failed %w",err)
	}
	log.Println("ver",ver,"method",method)

	_,err = conn.Write([]byte{socks5Ver,0x00})
	if err !=nil{
		return fmt.Errorf("write failed:%w",err)
	}
	return nil
}
