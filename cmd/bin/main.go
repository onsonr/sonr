package main

import (
	"github.com/sonr-io/core/cmd/lib"
	"github.com/sonr-io/core/internal/common"
	"github.com/sonr-io/core/internal/node"
	"google.golang.org/protobuf/proto"
)

func main() {
	req := &node.InitializeRequest{
		Profile: &common.Profile{
			FirstName: "Test",
			LastName:  "User",
			SName:     "TUser",
		},
	}
	buf, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	lib.Start(buf)
	select {}
}
