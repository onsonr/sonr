package highway

import (
	context "context"
	"fmt"

	rtv1 "github.com/sonr-io/blockchain/x/registry/types"
	rt "go.buf.build/grpc/go/sonr-io/sonr/registry"
)

// AccessName accesses a name.
func (s *HighwayServer) AccessName(ctx context.Context, req *rt.MsgAccessName) (*rt.MsgAccessNameResponse, error) {
	// // instantiate a query client for your `blog` blockchain
	// queryClient := registry.NewQueryClient(s.cosmos.Context)

	// // query the blockchain using the client's `didAll` method to get all dids
	// // store all dids in queryResp
	// queryResp, err := queryClient.Dids(context.Background(), &types.QueryDidsRequest{})
	// if err != nil {
	// 	return nil, err
	// }

	// print response from querying all the dids
	fmt.Print("\n\nAll Dids:\n\n")
	// fmt.Println(queryResp)
	return nil, ErrMethodUnimplemented
}

// RegisterName registers a name.
func (s *HighwayServer) RegisterName(ctx context.Context, req *rt.MsgRegisterName) (*rt.MsgRegisterNameResponse, error) {
	// account `alice` was initialized during `starport chain serve`
	accountName := "alice"

	// get account from the keyring by account name and return a bech32 address
	address, err := s.cosmos.Address(accountName)
	if err != nil {
		return nil, err
	}

	// define a message to create a did
	msg := &rtv1.MsgRegisterName{
		Creator:        address.String(),
		NameToRegister: req.NameToRegister,
		Payload:        req.Payload,
		//PublicKeyBuffer: req.GetJwk(),
	}

	// broadcast a transaction from account `alice` with the message to create a did
	// store response in txResp
	txResp, err := s.cosmos.BroadcastTx(accountName, msg)
	if err != nil {
		return nil, err
	}

	// print response from broadcasting a transaction
	logger.Infof("\n\nBroadcast Tx:\n\n%s\n\n", txResp)

	// fmt.Println(txResp)
	return nil, nil
}

// UpdateName updates a name.
func (s *HighwayServer) UpdateName(ctx context.Context, req *rt.MsgUpdateName) (*rt.MsgUpdateNameResponse, error) {
	return nil, ErrMethodUnimplemented
}
