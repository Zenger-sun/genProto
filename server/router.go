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

func (r *Router) RegisterRouter(msgType pb.MsgType, handler Handler) {
	r.Handler[msgType] = handler
}

func NewRouter() *Router {
	return &Router{Handler: make(map[pb.MsgType]Handler)}
}
