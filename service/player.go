package service

import (
	"errors"
	"genProto/model"
	"genProto/msg/pb"
	"genProto/server"
)

type Player struct {
	*model.DataPlayer
}

func (s *Svc) Login(packet *server.Packet, ctx *server.Context) error {
	req := packet.Msg.(*pb.LoginReq)
	if req == nil {
		return errors.New("decode msg err")
	}

	player := &Player{
		&model.DataPlayer{
			Id:        req.UserId,
			UserNanme: "test",
			Avatar:    0,
		},
	}
	s.NewPlayer(player)

	res := &pb.LoginRes{
		UserId: player.Id,
		Result: true,
	}

	server.Response(packet.Conn, pb.MsgType_MSG_LOGIN_RES, res)
	return nil
}
