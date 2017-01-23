package main

import (
	"encoding/json"
	"fmt"
	"gotcp/model"
	"net"
	"os"
	"strconv"
)

var ID int
var isbreak bool

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:9099")
	// go run test**.go 11
	args := os.Args
	ID, _ = strconv.Atoi(args[1])
	//TOID, _ := strconv.Atoi(args[2])
	register := model.Data{1, ID, ""}

	msg, _ := json.Marshal(register)
	json.Unmarshal(msg, register)

	fmt.Println(string(msg), register)
	conn.Write(msg)
	go clienthandleW(conn)
	go clienthandleR(conn)
	for {
		if isbreak {
			break
		}
	}
}

func clienthandleW(c net.Conn) {
	for {
		input := ""
		fmt.Scanf("%v", &input)
		chat := model.Data{2, ID, input}
		chatxt, _ := json.Marshal(chat)
		_, err := c.Write(chatxt)
		if err != nil {
			fmt.Println("error_read:", err)
			isbreak = true
		}
	}
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
