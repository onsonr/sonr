package bind

import (
	ath "github.com/sonr-io/core/internal/auth"
	sc "github.com/sonr-io/core/pkg/client"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ** ─── Node Type Checkers ────────────────────────────────────────────────────────
// Check type for Auth
func (mn *Node) isAuthType() bool {
	return mn.reqType == md.ConnectionRequest_AUTH
}

// Check Type for connect
func (mn *Node) isConnectType() bool {
	return mn.reqType == md.ConnectionRequest_CONNECT
}

// ** ─── Node Checkers ────────────────────────────────────────────────────────
func (mn *Node) isReady() bool {
	return mn.user.IsNotStatus(md.Status_STANDBY) || mn.user.IsNotStatus(md.Status_FAILED) || mn.isAuthType()
}

func (mn *Node) setConnected(val bool) {
	// Update Status
	su := mn.user.SetConnected(val)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		mn.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
		return
	}
	mn.call.OnStatus(data)
}

func (mn *Node) setAvailable(val bool) {
	// Update Status
	su := mn.user.SetAvailable(val)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		mn.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
		return
	}
	mn.call.OnStatus(data)
}

func (mn *Node) setStatus(newStatus md.Status) {
	// Set Status
	su := mn.user.SetStatus(newStatus)

	// Callback Status
	data, err := proto.Marshal(su)
	if err != nil {
		mn.handleError(md.NewError(err, md.ErrorMessage_MARSHAL))
		return
	}
	mn.call.OnStatus(data)
}

// ** ─── Node Initializers ────────────────────────────────────────────────────────
func (mn *Node) initialize(req *md.ConnectionRequest) {
	// Set Type
	mn.reqType = req.GetType()

	// Create Store - Start Auth Service
	if s, err := md.InitStore(req.GetDevice()); err == nil {
		mn.store = s
	}

	// Check Request Type
	if req.IsAuth() {
		// Start Auth Service
		mn.auth = ath.NewAuthService(req, mn.store, mn.callbackNode())
	} else {
		// Create User
		if u, err := md.NewUser(req, mn.store); err == nil {
			mn.user = u
		}

		// Create Client
		mn.client = sc.NewClient(mn.ctx, mn.user, mn.callbackNode())
	}
}
