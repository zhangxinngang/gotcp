package main

import (
	"encoding/json"
	"fmt"
	"gotcp/model"
	"net"
)

func main() {
	lis, _ := net.Listen("tcp", ":9099")
	for {
		conn, _ := lis.Accept()
		if conn != nil {
			fmt.Println(conn.LocalAddr())
			go gohandle(conn)
		}
	}
}

func gohandle(c net.Conn) {
	for {
		da := make([]byte, 128)
		n, err := c.Read(da)
		fmt.Println(string(da))
		data := model.Data{}
		err = json.Unmarshal(da[:n], &data)
		if err != nil {
			fmt.Println(err)
		}
		switch data.Type {
		case 1:
			fmt.Println("ssdfds")
			register(data.Id, c)
		case 2:
			broadcast(data.Id, c, data)
		}
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}

var chatConnMap = map[int]net.Conn{}

func register(id int, c net.Conn) {
	if chatConnMap[id] == nil {
		chatConnMap[id] = c
	}
	fmt.Println(chatConnMap)
	c.Write([]byte("req ok"))
}

func broadcast(id int, c net.Conn, data model.Data) {
	for i, c := range chatConnMap {
		if i == id {
			continue
		}
		c.Write([]byte(data.Context))
	}
}
