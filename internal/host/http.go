package host

import (
	"fmt"
	"net/http"

	"github.com/libp2p/go-libp2p-core/peer"
	gostream "github.com/libp2p/go-libp2p-gostream"
	p2phttp "github.com/libp2p/go-libp2p-http"
	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
)

// ** ─── REST HTTP Interfaces ────────────────────────────────────────────────────────
type HTTPHandler func(http.ResponseWriter, *http.Request)

type RestMethod interface {
	ID() peer.ID
	Name() string
	Handler() HTTPHandler
}

type restMethod struct {
	RestMethod
	id      peer.ID
	name    string
	handler HTTPHandler
}

// @ Method Returns Host ID
func (r *restMethod) ID() peer.ID {
	return r.id
}

// @ Method Returns HTTPMethod Name
func (r *restMethod) Name() string {
	return r.name
}

// @ MEthod Returns HTTPHandler Function
func (r *restMethod) Handler() HTTPHandler {
	return r.handler
}

// ^ Register Host with HTTP Method
func (h *HostNode) RegisterHTTPMethod(methodName string, handler HTTPHandler) (RestMethod, error) {
	// Initialize Listener
	listener, err := gostream.Listen(h.Host, p2phttp.DefaultP2PProtocol)
	if err != nil {
		return nil, err
	}

	// Create Method
	method := &restMethod{
		id:      h.Host.ID(),
		name:    methodName,
		handler: handler,
	}

	// Defer Listener
	defer listener.Close()

	// Start Go Fucntion
	go func(rm RestMethod) {
		http.HandleFunc(rm.Name(), rm.Handler())
		server := &http.Server{}
		server.Serve(listener)
	}(method)
	return method, nil
}

// ^ Utilize Host to Call HTTP Method
func (h *HostNode) CallHTTPMethod(methodName string, id peer.ID) (*http.Response, error) {
	// Register Transport
	tr := &http.Transport{}
	tr.RegisterProtocol("libp2p", p2phttp.NewTransport(h.Host))

	// Create Client
	client := &http.Client{Transport: tr}

	// Call Method
	return client.Get(fmt.Sprintf("libp2p://%s/%s", id.String(), methodName))
}

// ** ─── REST HTTP Structs ────────────────────────────────────────────────────────
// ^ Creates New HTTP Handler
func NewHTTPHandler(callback *md.NodeCallback) HTTPHandler {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Create Protobuf
		req := md.NewRestRequest(r)

		// Marshal Data
		buf, err := proto.Marshal(req)
		if err != nil {
			callback.Error(md.NewError(err, md.ErrorMessage_UNMARSHAL))
		}

		// Callback Request
		callback.APIRequest(buf)
	}
}
