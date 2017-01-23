package main

import (
	"fmt"
	"gotcp/model"
	"net"
	"time"
)

var ID int
var isbreak bool

func main() {
	client1 := model.RegisterCli(10)
	client2 := model.RegisterCli(11)

	go client1.Read()
	go client2.Read()

	chat := `{"type":3,"register":{"id":0},"broadcast":{"id":0,"context":""},"chat":{"id":10,"toid":11,"context":"aaaaaaaaaaa"}}`

	time.Sleep(time.Second * 11)

	client1.Write(chat)

	fmt.Println(client1, client2)

	client2.Write(`{"type":3,"chat":{"id":11,"toid":10,"context":"bbbbbbbbb"}}`)

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
