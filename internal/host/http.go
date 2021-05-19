package host

import (
	"net/http"

	gostream "github.com/libp2p/go-libp2p-gostream"
	p2phttp "github.com/libp2p/go-libp2p-http"
)

func (h *HostNode) HandleHTTP() {
	listener, _ := gostream.Listen(h.Host, p2phttp.DefaultP2PProtocol)
	defer listener.Close()
	go func() {
		http.HandleFunc("/link-request", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hi!"))
		})
		server := &http.Server{}
		server.Serve(listener)
	}()
}
