package server

import (
	"genProto/msg/pb"
	"log/slog"
	"net"

	"google.golang.org/protobuf/proto"
)

func listen(ctx *Context) error {
	laddr, err := net.ResolveTCPAddr("tcp", ctx.Addr)
	if err != nil {
		return err
	}

	l, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go readTCP(conn, ctx)
	}
}

func readTCP(conn net.Conn, ctx *Context) {
	buff := make([]byte, PACK_MAX_LEN)

	for {
		_, err := conn.Read(buff)
		switch err {
		case nil:
		default:
			slog.Warn("readTCP open conn err: ", err)
			return
		}

		pack, err := UnpackMsg(buff)
		if err != nil {
			slog.Warn("readTCP unpack head err: ", err)
			continue
		}
		pack.Conn = conn

		ctx.Router.RLock()
		handler, ok := ctx.Router.Handler[pb.MsgType(pack.MsgType)]
		ctx.Router.RUnlock()
		if !ok {
			continue
		}

		err = handler(pack, ctx)
		if err != nil {
			slog.Warn("handler err: ", err)
			continue
		}
	}
}

func Response(conn net.Conn, res pb.MsgType, msg proto.Message) {
	h := &Head{
		Len:     0,
		MsgType: uint16(res),
	}

	conn.Write(PackMsg(h, msg))
}
