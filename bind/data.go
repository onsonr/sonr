package sonr

import (
	"fmt"

	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// *************************** //
// *** Util to Convert Data ** //
// *************************** //
// ^  Provide bytes and converts to Proto ^  //
func (sn *Node) BytesToProto(bytes []byte, msg protoreflect.ProtoMessage) {
	// @ Source is bytes
	if bytes != nil {
		err := proto.Unmarshal(bytes, msg)
		if err != nil {
			sn.NewError(err, 4, pb.Error_PROTO, msg)
		}
	}
}

// ^  Provide json and converts to Proto ^  //
func (sn *Node) JsonToProto(json string, msg protoreflect.ProtoMessage) {
	err := protojson.Unmarshal([]byte(json), msg)
	if err != nil {
		sn.NewError(err, 4, pb.Error_PROTO, msg)
	}
}

// ^  Provide proto and converts to bytes ^  //
func (sn *Node) ProtoToBytes(msg protoreflect.ProtoMessage) []byte {
	bytes, err := proto.Marshal(msg)
	if err != nil {
		sn.NewError(err, 4, pb.Error_BYTES, msg)
	}
	return bytes
}

// ^  Provide proto and converts to json ^  //
func (sn *Node) ProtoToJson(msg protoreflect.ProtoMessage) string {
	// @ Source is bytes
	json, err := protojson.Marshal(msg)
	if err != nil {
		sn.NewError(err, 4, pb.Error_JSON)
	}
	return string(fmt.Sprintf("%s\n", string(json)))
}
