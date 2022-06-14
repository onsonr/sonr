package motor

import (
	"github.com/sonr-io/sonr/pkg/crypto"
	_ "golang.org/x/mobile/bind"
)

// Start starts the host, node, and rpc service.
func Start(reqBuf []byte) {
	_, err := crypto.Generate()
	if err != nil {
		panic(err)
	}
	//config := config.DefaultConfig(config.Role_MOTOR)
	// _, err := host.NewWasmHost(context.Background(), config)
	// if err != nil {
	// 	//golog.Fatal(err)
	// 	panic(err)
	// }
	//	node.Start(reqBuf)
	// Unmarshal request
	// req := &motor.InitializeRequest{}
	// if err := proto.Unmarshal(reqBuf, req); err != nil {
	// 	golog.Warn("%s - Failed to unmarshal InitializeRequest. Using defaults...", err)
	// }

	// Start the app
	//node.NewHighway(context.Background(), node.WithLogLevel(node.DebugLevel))
}

// Pause pauses the host, node, and rpc service.
func Pause() {
	//	node.Pause()
}

// Resume resumes the host, node, and rpc service.
func Resume() {
	//	node.Resume()
}

// Stop closes the host, node, and rpc service.
func Stop() {
	//	node.Exit(0)
}

// // Returns the cosmos compatible address of the given party.
// func (w *MPCWallet) AccountAddress(id ...party.ID) (types.AccAddress, error) {
// 	// c := types.GetConfig()
// 	// c.SetBech32PrefixForAccount("snr", "pub")
// 	bechAddr, err := w.Bech32Address(id...)
// 	if err != nil {
// 		return nil, err
// 	}
// 	acc, err := types.AccAddressFromBech32(bechAddr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return acc, nil
// }
