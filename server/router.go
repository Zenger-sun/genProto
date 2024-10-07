package server

import (
	"sync"

	"genProto/msg/pb"
)

type Handler func(packet *Packet, ctx *Context) error

type Router struct {
	sync.RWMutex
	Handler map[pb.MsgType]Handler
}
