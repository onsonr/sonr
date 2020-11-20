package sonr

import (
	"fmt"
	"reflect"

	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ NewError Generates error and returns to frontend and logs, also handles panic errors ^
func (sn *Node) NewError(err error, level pb.Error_Level, kind pb.Error_Kind, opts ...interface{}) error {
	// @ Initialize Default Message
	errMsg := fmt.Sprintf("⚡️⚠️  ERROR= %s based error %s severity => [ %s ]", kind.String(), level.String(), err)

	// @ Check Headers and Create Message
	if len(opts) > 1 {
		// Get Variable Info
		variable := opts[0]
		varType := reflect.TypeOf(variable)
		varName := varType.Name()

		// Set Message
		errMsg = fmt.Sprintf("⚡️⚠️  ERROR on [%s:%s]: %s based error %s severity => [ %s ]", varName, varType, kind.String(), level.String(), err)
	}

	// @ Create Protobuf Message
	errProto := &pb.Error{
		Kind:    kind,
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
