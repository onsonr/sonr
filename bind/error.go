package sonr

import (
	"fmt"

	pb "github.com/sonr-io/core/pkg/models"
)

// ^ NewError Generates error and returns to frontend and logs, also handles panic errors ^
func LogError(err error, code int32, kind pb.Error_Kind) *pb.Error {
	// @ Initialize Default Message
	level := pb.Error_Level(code)
	errMsg := fmt.Sprintf("⚡️⚠️  [ERROR]= %s based error %s severity => [ %s ]", kind.String(), level.String(), err)

	// @ Present Error
	fmt.Println(errMsg)

	// @ Create Protobuf Message
	return &pb.Error{
		Kind:    kind,
		Level:   level,
		Message: errMsg,
	}
}
