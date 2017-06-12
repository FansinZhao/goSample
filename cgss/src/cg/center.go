package cg

import (
	"encoding/json"
	"errors"
	"ipc"
	"sync"
)

var _ ipc.Server = &CenterServer{} //确认实现接口
type Message struct {
	From    string "from"
	To      string "to"
	Content string "content"
}

type Room struct {
}

type CenterServer struct {
	servers map[string]ipc.Server
	players []*Player
	rooms   []*Room
	mutext  sync.RWMutex
}

func NewCenterServer() *CenterServer {
	servers := make(map[string]ipc.Server)
	players := make([]*Player, 0)
	return &CenterServer{servers: servers, players: players}
}

func (server *CenterServer) addPlayer(param string) error {
	player := NewPlayer()
	err := json.Unmarshal([]byte(param), &player)
	if err != nil {
		return nil
	}
	server.mutext.Lock()
	defer server.mutext.Unlock()
	server.players = append(server.players, player)
	return nil
}

func (server *CenterServer) removePlayer(param string) error {
	server.mutext.Lock()
	defer server.mutext.Unlock()
	for i, v := range server.players {
		if v.Name == param {
			if len(server.players) == 1 {
				server.players = make([]*Player, 0)
			} else if i == len(server.players)-1 {
				server.players = server.players[:i-1]
			} else if i == 0 {
				server.players = server.players[1:]
			} else {
				server.players = append(server.players[:i-1], server.players[i+1:]...)
			}
			return nil
		}
	}
	return errors.New("没有该玩家")
}

func (server *CenterServer) listPlayers() (players string, err error) {
	server.mutext.Lock()
	defer server.mutext.Unlock()
	if len(server.players) > 0 {
		b, _ := json.Marshal(server.players)
		players = string(b)
	} else {
		err = errors.New("不存在任何玩家")
	}
	return
}

func (server *CenterServer) broadcast(param string) error {
	var msg Message
	err := json.Unmarshal([]byte(param), &msg)
	if err != nil {
		return err
	}

	server.mutext.Lock()
	defer server.mutext.Unlock()

	if len(server.players) > 0 {
		for _, p := range server.players {
			p.mq <- &msg
		}
	} else {
		err = errors.New("没有玩家需要通知")
	}
	return err
}

func (server *CenterServer) Handle(method, param string) *ipc.Response {
	switch method {
	case "addplayer":
		err := server.addPlayer(param)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "removeplayer":
		err := server.removePlayer(param)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "listplayer":
		players, err := server.listPlayers()
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{"200", players}
	case "broadcast":
		err := server.broadcast(param)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{Code: "200"}
	default:
		return &ipc.Response{Code: "404", Body: method + ":" + param}
	}
	return &ipc.Response{Code: "200"}
}

func (server *CenterServer) Name() string {
	return "CenterServer"
}
