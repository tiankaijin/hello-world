package proto

import "chat/model"

type Message struct {
	Cmd  string `json:"cmd"`
	Data string `json:"data"`
}

type LoginCmd struct {
	Id     int    `json:"user_id"`
	Passwd string `json:"passwd"`
}
type RegisterCmd struct {
	User model.User `json:"user"`
}

type LoginCmdRes struct {
	Code  int    `json:"code"`
	User  []int  `json:"users"`
	Error string `json:"error"`
}
type RegisterCmdRes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type UserListReq struct {
	Code  int    `json:"code"`
	Error string `jsong:"error"`
}

type UserListRes struct {
	Id []int `json:'"users"`
}
