package cg

import (
	"encoding/json"
	"errors"
	"ipc"
)

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterClient) AddPlayer(player *Player) error {
	b, err := json.Marshal(player)
	if err != nil {
		return err
	}

	resp, e := client.Call("addplayer", string(b))
	if e == nil && resp.Code == "200" {
		return nil
	}
	return err
}

func (client *CenterClient) RemovePlayer(name string) error {
	resp, e := client.Call("removeplayer", name)
	if e == nil && resp.Code == "200" {
		return nil
	}
	return e
}
func (client *CenterClient) ListPlayer(param string) (ps []*Player, err error) {
	resp, _ := client.Call("listplayer", param)
	if resp.Code != "200" {
		err = errors.New(resp.Code)
		return
	}
	err = json.Unmarshal([]byte(resp.Body), &ps)
	return
}

func (client *CenterClient) BroadCast(msg string) error {
	m := &Message{Content: msg}
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	resp, _ := client.Call("broadcast", string(b))
	if resp.Code == "200" {
		return nil
	}
	return errors.New(resp.Code)
}
