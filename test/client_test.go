package test

import (
	"fmt"
	"net"
	"testing"

	"genProto/msg/pb"
	"genProto/server"
)

func TestClinet(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return
	}
	defer conn.Close()

	head := &server.Head{
		Len:     0,
		MsgType: uint16(pb.MsgType_MSG_LOGIN_REQ),
	}
	msg := &pb.LoginReq{UserId: 1111}

	_, err = conn.Write(server.PackMsg(head, msg))
	if err != nil {
		return
	}

	pack := make([]byte, 1024)
	for {
		_, err := conn.Read(pack)
		if err != nil {
			continue
		}

		head, msg, err := server.UnpackMsg(pack)
		if err != nil {
			return
		}

		res := msg.(*pb.LoginRes)
		if res == nil {
			return
		}

		fmt.Println(head.MsgType, res.UserId, res.Result)
		break
	}
}

func BenchmarkClient(b *testing.B) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return
	}
	defer conn.Close()

	head := &server.Head{
		Len:     0,
		MsgType: uint16(pb.MsgType_MSG_LOGIN_REQ),
	}
	msg := &pb.LoginReq{UserId: 1111}

	for i := 0; i < b.N; i++ {
		conn.Write(server.PackMsg(head, msg))
	}
}