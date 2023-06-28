package handler

import (
	// "context"
	// "encoding/hex"

	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/gin-gonic/gin"
	// mdw "github.com/sonrhq/core/pkg/highway/middleware"
	// "google.golang.org/grpc"
)

type BroadcastTxResponse = txtypes.BroadcastTxResponse

// BroadcastSonrTx broadcasts a transaction on the Sonr blockchain network.
func BroadcastSonrTx(c *gin.Context) {

	// // Decode the hex-encoded transaction.
	// txRawBytes, err := hex.DecodeString(req.Tx)
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// }

	// // Create a connection to the gRPC server.
	// grpcConn, err := grpc.Dial(
	// 	mdw.GrpcEndpoint(),    // Or your gRPC server address.
	// 	grpc.WithInsecure(), // The Cosmos SDK doesn't support any transport security mechanism.
	// )
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// }
	// defer grpcConn.Close()


	// // Broadcast the tx via gRPC. We create a new client for the Protobuf Tx
	// // service.
	// txClient := txtypes.NewServiceClient(grpcConn)

	// // We then call the BroadcastTx method on this client.
	// grpcRes, err := txClient.BroadcastTx(
	// 	context.Background(),
	// 	&txtypes.BroadcastTxRequest{
	// 		Mode:    txtypes.BroadcastMode_BROADCAST_MODE_SYNC,
	// 		TxBytes: txRawBytes, // Proto-binary of the signed transaction, see previous step.
	// 	},
	// )
	// if err != nil {
	// 	c.JSON(500, gin.H{"error": err.Error()})
	// }
	// c.JSON(200, gin.H{"response": grpcRes})
}

// GetSonrTx returns a transaction on the Sonr blockchain network.
func GetSonrTx(c *gin.Context) {

}
