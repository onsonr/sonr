package main

import (
	"context"
	"errors"
	"log"
	"net"

	apir "github.com/sonr-io/sonr/pkg/motor"
	"github.com/sonr-io/sonr/third_party/types/common"
	api "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	rt "github.com/sonr-io/sonr/x/registry/types"
	_ "golang.org/x/mobile/bind"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	errWalletNotExists = errors.New("mpc wallet does not exist")
)

var (
	server   *MotorServer
	instance apir.MotorNode
	callback MotorCallback
)

type MotorCallback interface {
	OnDiscover(data []byte)
	OnWalletEvent(msg string, isDone bool)
}

type MotorServer struct {
	*grpc.Server
	api.MotorServiceServer
	MotorCallback
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server = &MotorServer{
		Server:        grpc.NewServer(),
		MotorCallback: common.DefaultCallback(),
	}
	// Register the Motor API
	api.RegisterMotorServiceServer(server, server)

	// Register reflection service on gRPC server.
	reflection.Register(server.Server)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *MotorServer) Initialize(ctx context.Context, req *api.InitializeRequest) (*api.InitializeResponse, error) {
	apir, err := apir.EmptyMotor(req, s)
	if err != nil {
		return nil, err
	}
	instance = apir
	callback = s

	// init objectBuilders
	//objectBuilders = make(map[string]*object.ObjectBuilder)

	// Return Initialization Response
	resp := api.InitializeResponse{
		Success: true,
	}

	if req.AuthInfo != nil {
		if res, err := instance.Login(api.LoginRequest{
			Did:       req.AuthInfo.Did,
			Password:  req.AuthInfo.Password,
			AesDscKey: req.AuthInfo.AesDscKey,
			AesPskKey: req.AuthInfo.AesPskKey,
		}); err == nil {
			return &api.InitializeResponse{
				Success: res.Success,
			}, nil
		}
	}
	return &resp, nil
}

func (s *MotorServer) CreateAccount(ctx context.Context, request *api.CreateAccountRequest) (*api.CreateAccountResponse, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	if res, err := instance.CreateAccount(*request); err == nil {
		go instance.Connect()
		return &res, nil
	} else {
		return nil, err
	}
}

func (s *MotorServer) Login(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	if res, err := instance.Login(*request); err == nil {
		go instance.Connect()
		return &res, nil
	} else {
		return nil, err
	}
}

// IssuePayment creates a send/receive token request to the specified address.
func (s *MotorServer) Payment(ctx context.Context, request *api.PaymentRequest) (*api.PaymentResponse, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	if res, err := instance.SendTokens(*request); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}

func (s *MotorServer) QuerySchema(ctx context.Context, request *api.QuerySchemaRequest) (*api.QueryWhatIsResponse, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	if request.Creator == "" && request.Did != "" {
		if instance == nil {
			return nil, errWalletNotExists
		}

		res, err := instance.QueryWhatIsByDid(request.Did)
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	if res, err := instance.QueryWhatIs(api.QueryWhatIsRequest{
		Creator: request.GetCreator(),
		Did:     request.GetDid(),
	}); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}

func (s *MotorServer) QueryWhatIsByCreator(ctx context.Context, request *api.QueryWhatIsByCreatorRequest) (*api.QueryWhatIsByCreatorResponse, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	if res, err := instance.QueryWhatIsByCreator(*request); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}

func (s *MotorServer) QueryWhereIs(ctx context.Context, request *api.QueryWhereIsRequest) (*api.QueryWhereIsResponse, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	res, err := instance.QueryWhereIs(*request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *MotorServer) CreateSchema(ctx context.Context, request *api.CreateSchemaRequest) (*api.CreateSchemaResponse, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}
	if res, err := instance.CreateSchema(*request); err == nil {
		return &res, nil
	} else {
		return nil, err
	}
}

func (s *MotorServer) QueryWhereIsByCreator(ctx context.Context, request *api.QueryWhereIsByCreatorRequest) (*api.QueryWhereIsByCreatorResponse, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	res, err := instance.QueryWhereIsByCreator(*request)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Stat returns general information about the Motor node its wallet and accompanying Account.
func (s *MotorServer) Stat(ctx context.Context, request *api.StatRequest) (*api.StatResponse, error) {
	if instance == nil {
		return nil, errWalletNotExists
	}

	doc := instance.GetDIDDocument()
	if doc == nil {
		return nil, errWalletNotExists
	}
	didDoc, err := rt.NewDIDDocumentFromPkg(doc)
	if err != nil {
		return nil, err
	}

	resp := api.StatResponse{
		Address:     instance.GetAddress(),
		Balance:     int32(instance.GetBalance()),
		DidDocument: didDoc,
	}
	return &resp, nil
}
