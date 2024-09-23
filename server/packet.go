package server

import (
	"bytes"
	"encoding/binary"
	"errors"

	"genProto/msg"
	"genProto/msg/pb"

	"google.golang.org/protobuf/proto"
)

const (
	PACK_MAX_LEN = 10240
	HEAD_LEN     = 6
)

type Head struct {
	Len       uint32
	MsgType   uint16
}

func UnpackMsg(data []byte) (*Head, proto.Message, error) {
	var head Head
	head.Len = uint32(data[0]) | uint32(data[1])<<8 | uint32(data[2])<<16 | uint32(data[3])<<24
	head.MsgType = uint16(data[4]) | uint16(data[5])<<8
	if head.Len > (PACK_MAX_LEN + HEAD_LEN) {
		return nil, nil, errors.New("msg len error!")
	}

	msgType := msg.GetMsgStruct(pb.MsgType(head.MsgType))
	err := proto.Unmarshal(data[HEAD_LEN:HEAD_LEN+head.Len], msgType)
	if err != nil {
		return nil, nil, err
	}

	return &head, msgType, nil
}

func PackMsg(head *Head, message proto.Message) []byte {
	msg, _ := proto.Marshal(message)
	head.Len = uint32(len(msg))

	buff := new(bytes.Buffer)
	binary.Write(buff, binary.LittleEndian, head.Len)
	binary.Write(buff, binary.LittleEndian, head.MsgType)
	buff.Write(msg)

	return buff.Bytes()
}
