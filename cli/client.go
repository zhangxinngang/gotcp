package main

import (
	"encoding/json"
	"fmt"
	"gotcp/model"
	"net"
	"strings"
)

type Client struct {
	Id   int
	Conn net.Conn
}

func Register(id int) Client {
	conn, _ := net.Dial("tcp", "127.0.0.1:9099")
	// go run test**.go 11
	//TOID, _ := strconv.Atoi(args[2])
	register := model.Data{Type: 1, Register: model.Register{Id: id}}

	msg, _ := json.Marshal(register)
	json.Unmarshal(msg, register)

	fmt.Println(string(msg), register)
	conn.Write(msg)

	conn.Read(msg)
	for !strings.Contains(string(msg), "resp ok") {
		conn.Read(msg)
	}

	return Client{
		Id:   id,
		Conn: conn,
	}
}

func (this *Client) Read() {
	for {
		b := make([]byte, 128)
		n, err := this.Conn.Read(b)
		if err != nil {
			fmt.Println("error_read:", err)
			break
		}
		if string(b) != "" {
			fmt.Println(this.Id, "receive", string(b[:n]))
		}
	}
}

func (this *Client) Write(input string) {
	_, err := this.Conn.Write([]byte(input))
	if err != nil {
		fmt.Println("error_read:", err)
		isbreak = true
	}
}
