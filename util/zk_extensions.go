/*
 * Copyright 2017 XLAB d.o.o.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package util

import (
	"google.golang.org/grpc"

	pb "go.buf.build/grpc/go/sonr-io/core/host/zk/v1"
)

// ServerStream is an interface that fits all the auto-generated server
// stream interfaces declared within this package.
type ServerStream interface {
	Send(*pb.Message) error
	Recv() (*pb.Message, error)
	grpc.ServerStream
}

// ClientStream is an interface that fits all the auto-generated client
// stream interfaces declared within this package.
type ClientStream interface {
	Send(*pb.Message) error
	Recv() (*pb.Message, error)
	grpc.ClientStream
}
