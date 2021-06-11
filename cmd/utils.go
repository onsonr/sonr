package bind

import (
	"context"

	"github.com/getsentry/sentry-go"
	sc "github.com/sonr-io/core/pkg/client"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// @ Return URLLink
func URLLink(url string) []byte {
	// Create Link
	link := md.NewURLLink(url)

	// Marshal
	bytes, err := proto.Marshal(link)
	if err != nil {
		return nil
	}
	return bytes
}

// @ Gets User from Storj
func Storj(data []byte) []byte {
	// Unmarshal Request
	request := &md.StorjRequest{}
	proto.Unmarshal(data, request)

	switch request.Data.(type) {
	// @ Put USER
	case *md.StorjRequest_User:
		// Put User
		err := sc.PutUser(context.Background(), request.StorjApiKey, request.GetUser())
		if err != nil {
			sentry.CaptureException(err)

			// Create Response
			resp := &md.StorjResponse{
				Data: &md.StorjResponse_Success{
					Success: false,
				},
			}

			// Marshal
			bytes, err := proto.Marshal(resp)
			if err != nil {
				return nil
			}
			return bytes
		}
		// Create Response
		resp := &md.StorjResponse{
			Data: &md.StorjResponse_Success{
				Success: true,
			},
		}

		// Marshal
		bytes, err := proto.Marshal(resp)
		if err != nil {
			return nil
		}
		return bytes

	// @ Get USER
	case *md.StorjRequest_Prefix:
		// Get User from Uplink
		user, err := sc.GetUser(context.Background(), request.StorjApiKey, request.GetPrefix())
		if err != nil {
			sentry.CaptureException(err)
			return nil
		}

		// Create Response
		resp := &md.StorjResponse{
			Data: &md.StorjResponse_User{
				User: user,
			},
		}

		// Marshal
		bytes, err := proto.Marshal(resp)
		if err != nil {
			return nil
		}
		return bytes
	}
	return nil
}
