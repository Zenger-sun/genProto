package msg

import (
	"genProto/msg/pb"

	"google.golang.org/protobuf/proto"
)

func GetMsgStruct(msgCode pb.MsgType) proto.Message {
	switch msgCode {
	case pb.MsgType_MSG_LOGIN_REQ:
		return &pb.LoginReq{}
	case pb.MsgType_MSG_LOGIN_RES:
		return &pb.LoginRes{}
	}

	return nil
}
