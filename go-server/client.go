package main

import (
	"fmt"
	"net"
	"time"
	"strconv"
)

const (
	addr = "127.0.0.1:3333"
)

func main() {

	for i := 0; i < 1; i++{
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			fmt.Println("连接服务器失败：", err.Error())
			return
		}

		fmt.Printf("客戶端 %d 已连接服务器 \n", i)
		defer conn.Close()
		Client(conn)
		time.Sleep(time.Second*2)
	}

	channel := make(chan int32)

	<-channel
}

func Client(conn net.Conn) {

	i := 0
	data := "I am a client msg-"

	msg_head := "Hello Go!"
	//msg_head := ""


	go func() {
		for {
			str := strconv.Itoa(i)
			msg_conten_byte := []byte(data + str)

			msg_size_byte := make([]byte, 2)
			msg_size_byte[0] = byte(len(msg_conten_byte) >> 8)
			msg_size_byte[1] = byte(len(msg_conten_byte) % 256)

			msg := msg_head + string(msg_size_byte[0]) + string(msg_size_byte[1]) + string(msg_conten_byte)

			conn.Write([]byte(msg))
			i++
			if i >= 126 {
				i = 33
			}
			time.Sleep(time.Second * 1)
		}

	}()

	go func() {

		for {
			buf := make([]byte, 1024)
			i, err := conn.Read(buf)
			if err != nil {
				fmt.Println("error:", err.Error())
			}

			if i > 2 {
				fmt.Printf("客户端 %d says ： %s \n", int(buf[0]) << 8 + int(buf[1]), string(buf[2:i]))
			} else {
				fmt.Printf("server sends  ： %s \n", string(buf[:i]))
			}

		}

	}()

}
