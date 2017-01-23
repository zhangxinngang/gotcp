package main

import (
	"fmt"
	"net"
	"time"
)

var ID int
var isbreak bool

func main() {
	client1 := Register(10)
	client2 := Register(11)

	go client1.Read()
	go client2.Read()

	chat := `{"type":3,"register":{"id":0},"broadcast":{"id":0,"context":""},"chat":{"id":10,"toid":11,"context":"aaaaaaaaaaa"}}`

	client1.Write(chat)

	fmt.Println(client1, client2)

	time.Sleep(time.Second * 1)

	client1.Conn.Close()
	client2.Conn.Close()
}

func ClinetW(c net.Conn) {
	clienthandleW(c)
}

func clienthandleW(c net.Conn) {
	for {
		input := ""
		fmt.Scanf("%v", &input)
		_, err := c.Write([]byte(input))
		if err != nil {
			fmt.Println("error_read:", err)
			isbreak = true
		}
	}
}

func ClinetR(c net.Conn) {
	clienthandleR(c)
}

func clienthandleR(c net.Conn) {
	for {
		b := make([]byte, 128)
		n, err := c.Read(b)
		if err != nil {
			fmt.Println("error_read:", err)
			isbreak = true
		}
		if string(b) != "" {
			fmt.Println(string(b[:n]))
		}
	}
}
