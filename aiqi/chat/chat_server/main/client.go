package main

import (
	"chat/proto"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

type Client struct {
	conn net.Conn
	buf  [8192]byte
}

func (p *Client) writePackage(data []byte) (err error) {
	packLen := uint32(len(data))
	binary.BigEndian.PutUint32(p.buf[0:4], packLen)
	n, err := p.conn.Write(p.buf[0:4])
	if err != nil || n != 4 {
		fmt.Println("write data len failed")
		return
	}
	n, err = p.conn.Write(data)
	if err != nil {
		fmt.Println("write data failed")
		return
	}
	if n != int(packLen) {
		fmt.Println("write data not finished")
		err = errors.New("write data not finished")
		return
	}
	return

}
func (p *Client) readPackage() (msg proto.Message, err error) {
	n, err := p.conn.Read(p.buf[0:4])
	if n != 4 {
		if err != io.EOF {
			err = errors.New("read header failed")
		}

		return
	}
	var packLen uint32
	packLen = binary.BigEndian.Uint32(p.buf[0:4])

	n, err = p.conn.Read(p.buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}
	err = json.Unmarshal(p.buf[0:packLen], &msg)

	if err != nil {
		fmt.Println("unmarshal error, ", err)
	}
	return
}
func (p *Client) Proccess() (err error) {
	for {
		var msg proto.Message
		msg, err = p.readPackage()
		if err != nil {
			return err
		}
		err = p.proccessMsg(msg)
		if err != nil {
			return err
		}
	}
	return
}
func (p *Client) proccessMsg(msg proto.Message) (err error) {

	switch msg.Cmd {
	case UserLogin:
		err = p.login(msg)
	case UserRegister:
		err = p.register(msg)
	default:
		err = errors.New("unsupport message")
	}
	return
}
func (p *Client) loginResp(err error) {
	var resMsg proto.Message
	resMsg.Cmd = UserLoginRes

	var loginRes proto.LoginCmdRes
	loginRes.Code = 200

	userMap := clientMgr.GetAllUsers()
	for userid, _ := range userMap {
		loginRes.User = append(loginRes.User, userid)
	}
	if err != nil {

		loginRes.Code = 500
		loginRes.Error = fmt.Sprintf("%v", err)
	}

	data, err := json.Marshal(loginRes)
	if err != nil {
		fmt.Println("marshal failed, ", err)
		return
	}
	resMsg.Data = string(data)
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("marshal failed, ", err)
		return
	}
	p.writePackage(data)

}
func (p *Client) registerResp(err error) {
	var resMsg Message
	resMsg.Cmd = UserRegisterRes

	var registerRes RegisterCmdRes
	registerRes.Code = 200
	if err != nil {

		registerRes.Code = 500
		registerRes.Error = fmt.Sprintf("%v", err)
	}

	data, err := json.Marshal(registerRes)
	if err != nil {
		fmt.Println("marshal failed, ", err)
		return
	}
	resMsg.Data = string(data)
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("marshal failed, ", err)
		return
	}
	p.writePackage(data)

}

func (p *Client) login(msg proto.Message) (err error) {

	var cmd proto.LoginCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}
	_, err = mgr.Login(cmd.Id, cmd.Passwd)
	defer p.loginResp(err)
	if err != nil {
		return
	}

	clientMgr.AddClient(cmd.Id, p)
	return
}
func (p *Client) register(msg proto.Message) (err error) {
	var cmd proto.RegisterCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}

	err = mgr.Register(&cmd.User)
	defer p.registerResp(err)
	if err != nil {
		return
	}
	return
}
