package service

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/sonr-io/core/pkg/util"
	"google.golang.org/api/option"
)

var isPushEnabled = false

type PushService struct {
	ctx       context.Context
	key       string
	app       *firebase.App
	client    *messaging.Client
	options   *md.ConnectionRequest_ServiceOptions
	pushToken string
}

// Returns New Push Client
func (sc *serviceClient) StartPush() *md.SonrError {
	// Initialize
	if sc.request.GetServiceOptions().GetPush() {
		// Logging
		md.LogActivate("Push Service")

		// Obtain a messaging.Client from the App.
		ctx := context.Background()
		opt := option.WithCredentialsFile(sc.request.GetApiKeys().GetPushKeyPath())
		config := &firebase.Config{ProjectID: util.FIRE_PROJECT_ID}

		// Create New Firebase Client
		app, err := firebase.NewApp(context.Background(), config, opt)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_PUSH_START_APP)
		}

		// Create New Push Client
		client, err := app.Messaging(ctx)
		if err != nil {
			return md.NewError(err, md.ErrorMessage_PUSH_START_MESSAGING)
		}

		// Return Push Interface
		sc.Push = &PushService{
			key:       sc.request.GetApiKeys().GetPushKeyPath(),
			app:       app,
			client:    client,
			ctx:       ctx,
			pushToken: sc.pushToken,
			options:   sc.request.GetServiceOptions(),
		}
		isPushEnabled = true
		md.LogSuccess("Push Notifications Activation")
	}
	return nil
}

// Push pushes a single message to a single Peer
func (pc *PushService) push(msg *md.PushMessage) *md.SonrError {
	// Check for Push Token
	pushToken, serr := msg.GetPeer().PushToken()
	if serr != nil {
		return serr
	}

	// Create Message
	pushMsg := &messaging.Message{
		Token: pushToken,
		Data:  msg.GetData(),
	}

	// Send Message
	result, err := pc.client.Send(pc.ctx, pushMsg)
	if err != nil {
		return md.NewError(err, md.ErrorMessage_PUSH_SINGLE)
	}

	// Logging
	md.LogSuccess("Pushed Message: " + result)
	return nil
}

// PushMulti pushes Multiple messages to list of Peers
func (pc *PushService) pushMulti(msg *md.PushMessage, peers []*md.Peer) *md.SonrError {
	// Initialize List of Tokens
	tokens := make([]string, len(peers))
	for i, peer := range peers {
		tokens[i] = peer.Id.GetPushToken()
	}

	// Create Message
	multiPushMsg := &messaging.MulticastMessage{
		Tokens: tokens,
		Data:   msg.GetData(),
	}

	// Send Message
	result, err := pc.client.SendMulticast(pc.ctx, multiPushMsg)
	if err != nil {
		return md.NewError(err, md.ErrorMessage_PUSH_SINGLE)
	}

	// Logging
	md.LogInfo(fmt.Sprintf("Succesful Push Count: %v \n Failed Push Count: %v", result.SuccessCount, result.FailureCount))
	return nil
}

// pushSelf method sends push notification to own device
func (pc *PushService) pushSelf(msg *md.PushMessage) *md.SonrError {
	// Create Message
	pushMsg := &messaging.Message{
		Token: pc.pushToken,
		Data:  msg.GetData(),
	}

	// Send Message
	result, err := pc.client.Send(pc.ctx, pushMsg)
	if err != nil {
		return md.NewError(err, md.ErrorMessage_PUSH_SINGLE)
	}

	// Logging
	md.LogSuccess("Pushed Message: " + result)
	return nil
}
