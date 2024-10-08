package test

import (
	"fmt"
	"genProto/msg/pb"
	"genProto/server"
	"net"
	"sync"
	"testing"
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

		pack, err := server.UnpackMsg(pack)
		if err != nil {
			return
		}

		res := pack.Msg.(*pb.LoginRes)
		if res == nil {
			return
		}

		fmt.Println(head.MsgType, res.UserId, res.Result)

		conn.Close()
		break
	}
}

func TestMultiClinet(t *testing.T) {
	var connList []net.Conn

	for i := 0; i < 10000; i++ {
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			return
		}

		connList = append(connList, conn)
	}

	var wg sync.WaitGroup
	for _, conn := range connList {
		conn := conn
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				conn.Close()
			}()

			head := &server.Head{
				Len:     0,
				MsgType: uint16(pb.MsgType_MSG_LOGIN_REQ),
			}
			msg := &pb.LoginReq{UserId: 1111}

			_, err := conn.Write(server.PackMsg(head, msg))
			if err != nil {
				return
			}

			pack := make([]byte, 1024)
			for {
				_, err := conn.Read(pack)
				if err != nil {
					continue
				}

				pack, err := server.UnpackMsg(pack)
				if err != nil {
					return
				}

				res := pack.Msg.(*pb.LoginRes)
				if res == nil {
					return
				}

				break
			}
		}()
	}
	wg.Wait()
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