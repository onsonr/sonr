package sonr

import (
	"fmt"
	"strings"

	pb "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ^ NewError Generates error and returns to frontend and logs, also handles panic errors ^
func (sn *Node) NewError(err error, level pb.Error_Level, kind pb.Error_Kind, headers ...string) error {
	// @ Initialize
	var errMsg string

	// @ Check Headers and Create Message
	if len(headers) == 1 {
		// Check Header type
		data := headers[0]

		// Variable Prefix
		if strings.Contains(data, "V=") {
			// Set Error Message
			errMsg = createVariableMessage(err, level, kind, data)
		} else {
			// Set Error Message
			errMsg = createMethodMessage(err, level, kind, data)
		}
	} else if len(headers) == 2 {
		// Get headers
		m := headers[0]
		v := headers[1]

		// Set Error Message
		errMsg = createFullMessage(err, level, kind, m, v)
	} else {
		errMsg = createDefaultMessage(err, level, kind)
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

// ** Error Message with Method **
func createMethodMessage(err error, lvl pb.Error_Level, k pb.Error_Kind, method string) string {
	return fmt.Sprintf("⚡️⚠️  %s(): %s based error occured of %s severity => [ %s ]", method, k.String(), lvl.String(), err)
}

// ** Error Message with Variable **
func createVariableMessage(err error, lvl pb.Error_Level, k pb.Error_Kind, variable string) string {
	return fmt.Sprintf("⚡️⚠️  [ERROR=%s]: %s based error occured of %s severity => [ %s ]", variable, k.String(), lvl.String(), err)
}

// ** Error Message with Variable and Method **
func createFullMessage(err error, lvl pb.Error_Level, k pb.Error_Kind, method string, variable string) string {
	return fmt.Sprintf("⚡️⚠️  %s(): %s based error occured of %s severity on %s => [ %s ]", method, k.String(), lvl.String(), variable, err)
}

// ** Error Message Default **
func createDefaultMessage(err error, lvl pb.Error_Level, k pb.Error_Kind) string {
	return fmt.Sprintf("⚡️⚠️  [ERROR]: %s based error occured of %s severity => [ %s ]", k.String(), lvl.String(), err)
}
