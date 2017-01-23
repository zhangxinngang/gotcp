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
		fmt.Println(string(da))
		data := model.Data{}
		err = json.Unmarshal(da[:n], &data)
		if err != nil {
			fmt.Println(err, da[:n])
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
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

var chatConnMap = map[int]net.Conn{}

func register(data model.Register, c net.Conn) {
	chatConnMap[data.Id] = c
	fmt.Println(chatConnMap)
	c.Write([]byte("resp ok"))
}

func broadcast(data model.Broadcast, c net.Conn) {
	for i, cn := range chatConnMap {
		if i == data.Id {
			continue
		}
		cn.Write([]byte(data.Context))
	}
	c.Write([]byte("resp ok"))
}

func chat(data model.Chat, c net.Conn) {
	tocn := chatConnMap[data.ToId]
	if tocn == nil {
		return
	}
	tocn.Write([]byte(data.Context))
	c.Write([]byte("resp ok"))
}
