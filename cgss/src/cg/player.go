package cg

import (
	"fmt"
)

type Player struct {
	Name  string "名称"
	Level int    "lever"
	Exp   int    "exp"

	mq chan *Message
}

func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{"p1", 0, 0, m}

	go func(p *Player) {
		for {
			msg := <-p.mq
			fmt.Println(p.Name, "接收到消息", msg.Content)
		}
	}(player)

	return player
}
