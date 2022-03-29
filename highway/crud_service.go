package highway

import (
	context "context"
	"fmt"

	rt "go.buf.build/grpc/go/sonr-io/sonr/registry"
)

// AccessService accesses a service.
func (s *HighwayServer) AccessService(ctx context.Context, req *rt.MsgAccessService) (*rt.MsgAccessServiceResponse, error) {
	// // instantiate a query client for your `blog` blockchain
	// queryClient := types.NewQueryClient(s.cosmos.Context)

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

// RegisterService registers a service.
func (s *HighwayServer) RegisterService(ctx context.Context, req *rt.MsgRegisterService) (*rt.MsgRegisterServiceResponse, error) {
	// // account `alice` was initialized during `starport chain serve`
	// accountName := "alice"

	// // get account from the keyring by account name and return a bech32 address
	// address, err := s.cosmos.Address(accountName)
	// if err != nil {
	// 	return nil, err
	// }

	// // define a message to create a did
	// msg := &types.MsgCreateDidDocument{
	// 	Creator: address.String(),
	// }

	// // broadcast a transaction from account `alice` with the message to create a did
	// // store response in txResp
	// txResp, err := s.cosmos.BroadcastTx(accountName, msg)
	// if err != nil {
	// 	return nil, err
	// }

	// print response from broadcasting a transaction
	fmt.Print("MsgCreateDidDocument:\n\n")
	// fmt.Println(txResp)
	return nil, ErrMethodUnimplemented
}

// UpdateService updates a service.
func (s *HighwayServer) UpdateService(ctx context.Context, req *rt.MsgUpdateService) (*rt.MsgUpdateServiceResponse, error) {
	return nil, ErrMethodUnimplemented
}
