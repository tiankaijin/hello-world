package main

import (
	"errors"
	"fmt"
)

var (
	clientMgr *ClientMgr
)

type ClientMgr struct {
	onlineUsers map[int]*Client
}

func init() {
	clientMgr = &ClientMgr{
		onlineUsers: make(map[int]*Client, 1024),
	}
}
func (p *ClientMgr) AddClient(userId int, client *Client) {
	p.onlineUsers[userId] = client
}

func (p *ClientMgr) GetClient(userId int) (client *Client, err error) {
	client, ok := p.onlineUsers[userId]
	if !ok {

		err = errors.New(fmt.Sprintf("user %d not exist", userId))
		return
	}
	return
}
func (p *ClientMgr) GetAllUsers() map[int]*Client {
	return p.onlineUsers
}
