package service

import (
	"errors"
	"fmt"
	"net"

	"genProto/msg/pb"
	"genProto/server"

	"google.golang.org/protobuf/proto"
)

func Login(msg proto.Message, conn net.Conn) error {
	req := msg.(*pb.LoginReq)
	if req == nil {
		return errors.New("decode msg err")
	}

	fmt.Println(req.UserId)

	res := &pb.LoginRes{
		UserId: req.UserId,
		Result: true,
	}

	h := &server.Head{
		Len:     0,
		MsgType: uint16(pb.MsgType_MSG_LOGIN_RES),
	}

	server.Response(conn, server.PackMsg(h, res))
	return nil
}
