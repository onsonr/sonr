package sonr

import (
	"fmt"
	"strings"

	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// NewError Generates error and returns to frontend and logs, also handles panic errors
func (sn *Node) NewError(err error, code int32, kind string, method string) error {
	// @ Generate Error Message
	level := pb.Error_ErrorLevel(code)
	errMsg := fmt.Sprintf("%s(): %s type of error occured of %s severity", method, kind, level.String())

	// @ Create Protobuf Message
	errProto := &pb.Error{
		Kind:    strings.ToUpper(kind),
		Level:   level,
		Message: errMsg,
	}

	// @ Convert to bytes
	errBytes, err := proto.Marshal(errProto)
	if err != nil {
		fmt.Println("Error Marshaling Error XD")
	}

	// @ Present Error
	sn.Call.OnError(errBytes)
	fmt.Println(errMsg)

	// @ Handle for Level Severity
	if level == pb.Error_CRITICAL {
		panic(err)
	}

	// @ Not Panic Error
	return err
}
