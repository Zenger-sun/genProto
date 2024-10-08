package service

import (
	"genProto/msg/pb"
	"genProto/server"
)

type Svc struct {
	*server.Context
	Player map[uint32]*Player
}

func (s *Svc) NewPlayer(p *Player) {
	s.Player[p.Id] = p
}

func (s *Svc) GetPlayer(uid uint32) *Player {
	if p, ok := s.Player[uid]; ok {
		return p
	}

	return nil
}

func (s *Svc) Init() {
	s.RegisterRouter(pb.MsgType_MSG_LOGIN_REQ, s.Login)
}

func NewSvc(ctx *server.Context) *Svc {
	return &Svc{
		Context: ctx,
		Player:  make(map[uint32]*Player),
	}
}
