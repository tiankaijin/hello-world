package main

import (
	"chat/model"
	"chat/proto"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

func writePackage(conn net.Conn, data []byte) (err error) {
	var buf [4]byte
	packLen := uint32(len(data))
	binary.BigEndian.PutUint32(buf[:], packLen)
	n, err := conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("write data len failed")
		return
	}
	n, err = conn.Write(data)
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
func readPackage(conn net.Conn) (msg proto.Message, err error) {
	var buf [8192]byte
	n, err := conn.Read(buf[0:4])
	if err != nil || n != 4 {
		err = errors.New("read header failed")
		return
	}
	var packLen uint32
	packLen = binary.BigEndian.Uint32(buf[0:4])

	n, err = conn.Read(buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}
	err = json.Unmarshal(buf[0:packLen], &msg)

	if err != nil {
		fmt.Println("error read last")
	}
	return
}
func login(conn net.Conn) (err error) {
	var msg proto.Message
	msg.Cmd = proto.UserLogin

	var loginCmd proto.LoginCmd
	loginCmd.Id = 2
	loginCmd.Passwd = "123456"

	data, err := json.Marshal(loginCmd)
	if err != nil {
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}
	writePackage(conn, data)
	msg, err = readPackage(conn)
	if err != nil {
		return
	}
	fmt.Println(msg)
	return

}
func register(conn net.Conn, user model.User) {
	var msg proto.Message
	msg.Cmd = proto.UserRegister

	var registerCmd proto.RegisterCmd
	registerCmd.User = user
	data, err := json.Marshal(registerCmd)
	if err != nil {
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}
	writePackage(conn, data)
	msg, err = readPackage(conn)
	if err != nil {
		return
	}
	fmt.Println(msg)

}

func main() {
	conn, err := net.Dial("tcp", "localhost:10000")
	if err != nil {
		fmt.Println("Error dialing, ", err.Error())
		return
	}
	defer conn.Close()
	err = login(conn)
	//user := model.GetUser()
	//register(conn, user)
	for {

	}
	if err != nil {
		fmt.Println("login failed,err:", err)
		return
	}
}
