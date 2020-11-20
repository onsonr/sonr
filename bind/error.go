package sonr

import (
	"fmt"
	"reflect"

	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ NewError Generates error and returns to frontend and logs, also handles panic errors ^
func (sn *Node) NewError(err error, code int32, kind pb.Error_Kind, checkVar ...interface{}) error {
	// @ Initialize Default Message
	var typeVar reflect.Type
	typeName := "ERROR"
	level := pb.Error_Level(code)
	errMsg := fmt.Sprintf("⚡️⚠️  [%s]= %s based error %s severity => [ %s ]", typeName, kind.String(), level.String(), err)

	// @ Check Vars
	if len(checkVar) > 0 {
		// Update Values
		typeVar = reflect.TypeOf(checkVar)
		typeName = typeVar.Name()

		// Update Message
		errMsg = fmt.Sprintf("⚡️⚠️  [%s:%s]= %s based error %s severity => [ %s ]", typeName, typeVar, kind.String(), level.String(), err)
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
