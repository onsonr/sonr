package crypto

import (
	"fmt"
	"sync"

	"github.com/taurusgroup/multi-party-sig/pkg/party"
	"github.com/taurusgroup/multi-party-sig/pkg/pool"
	"github.com/taurusgroup/multi-party-sig/pkg/protocol"
	"github.com/taurusgroup/multi-party-sig/protocols/frost"
)

func Generate(bio1 string) error {
	participants := []party.ID{party.ID(bio1), "vault", "shared"}
	threshold := 1
	net := NewNetwork(participants)

	pl := pool.NewPool(0)
	defer pl.TearDown()

	var wg sync.WaitGroup
	for _, id := range participants {
		wg.Add(1)
		go func(id party.ID) {
			pl := pool.NewPool(0)
			defer pl.TearDown()
			// if err := All(id, ids, threshold, messageToSign, net, &wg, pl); err != nil {
			if conf, err := frostKeygenTaproot(id, participants, threshold, net, &wg); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("success")
				fmt.Printf("%+v\n", conf)
			}
		}(id)
	}
	wg.Wait()
	fmt.Println("done.")

	return nil
}

func frostKeygenTaproot(id party.ID, ids party.IDSlice, threshold int, n *Network, wg *sync.WaitGroup) (*frost.TaprootConfig, error) {
	defer wg.Done()

	h, err := protocol.NewMultiHandler(frost.KeygenTaproot(id, ids, threshold), nil)
	if err != nil {
		return nil, err
	}
	handlerLoop(id, h, n)
	r, err := h.Result()
	if err != nil {
		return nil, err
	}

	return r.(*frost.TaprootConfig), nil
}

func handlerLoop(id party.ID, h protocol.Handler, network *Network) {
	for {
		fmt.Println("for")
		select {

		// outgoing messages
		case msg, ok := <-h.Listen():
			fmt.Println("listen")
			fmt.Println(msg)
			if !ok {
				<-network.Done(id)
				// the channel was closed, indicating that the protocol is done executing.
				return
			}
			go network.Send(msg)

			// incoming messages
		case msg := <-network.Next(id):
			fmt.Println("next")
			fmt.Println(msg)
			h.Accept(msg)
		}
	}
}
