package model

const (
	TYPE_REGISTER  = 1
	TYPE_BROADCAST = 2
	TYPE_CHAT      = 3
)

type Data struct {
	Type      int `json:"type"`
	Register  `json:"register,omitempty"`
	Broadcast `json:"broadcast,omitempty"`
	Chat      `json:"chat,omitempty"`
}

/*
{"type":3,"register":{"id":0},"broadcast":{"id":0,"context":""},"chat":{"id":10,"toid":11,"context":"aaaaaaaaaaa"}}


{"type":3,"chat":{"id":10,"toid":11,"context":"helloworld"}}

*/

type Register struct {
	Id int `json:"id"`
}

type Broadcast struct {
	Id      int    `json:"id"`
	Context string `json:"context"`
}

type Chat struct {
	Id      int    `json:"id"`
	ToId    int    `json:"toid"`
	Context string `json:"context"`
}
