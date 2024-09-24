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

	server.Response(conn, pb.MsgType_MSG_LOGIN_RES, res)
	return nil
}
