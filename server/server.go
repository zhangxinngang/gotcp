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
		if err != nil {

			return
		}
		data := model.Data{}
		err = json.Unmarshal(da[:n], &data)
		if err != nil {
			return
		}
		switch data.Type {
		case model.TYPE_REGISTER:
			register(data.Register, c)
		case model.TYPE_BROADCAST:
			broadcast(data.Broadcast, c)
		case model.TYPE_CHAT:
			chat(data.Chat, c)
		}
	}
}

var chatCliMap = map[int]model.Client{}

func register(data model.Register, c net.Conn) {
	cli := model.Client{Id: data.Id, Conn: c}
	chatCliMap[data.Id] = cli
	fmt.Println(chatCliMap)
	c.Write([]byte("resp ok"))
}

func broadcast(data model.Broadcast, c net.Conn) {
	for _, cli := range chatCliMap {
		if cli.Id == data.Id {
			continue
		}
		cli.Write(data.Context)
	}
	c.Write([]byte("resp ok"))
}

func chat(data model.Chat, c net.Conn) {
	cli := chatCliMap[data.ToId]
	if cli.Id == 0 {
		return
	}
	cli.Write(data.Context)
	c.Write([]byte("resp ok"))
}
