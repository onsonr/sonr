/*
 * Copyright 2017 XLAB d.o.o.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package server

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"math"
	"math/big"
	"net"

	"net/http"

	"github.com/go-redis/redis"
	"github.com/golang/protobuf/ptypes/empty"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sonr-io/core/config"
	"github.com/sonr-io/core/crypto/cl"
	"github.com/sonr-io/core/crypto/ec"
	"github.com/sonr-io/core/crypto/ecpseudsys"
	"github.com/sonr-io/core/crypto/ecschnorr"
	"github.com/sonr-io/core/crypto/pseudsys"
	"github.com/sonr-io/core/crypto/schnorr"
	zk "github.com/sonr-io/core/host/zk"
	"github.com/sonr-io/core/log"
	pb "go.buf.build/grpc/go/sonr-io/core/host/zk/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

const (
	// Curve to be used in all schemes using elliptic curve arithmetic.
	curve = ec.P256
)

// EmmyServer is an interface composed of all the auto-generated server interfaces that
// declare gRPC handler functions for emmy protocols and schemes.
type EmmyServer interface {
	pb.PseudonymSystemServer
	pb.PseudonymSystemCAServer
	pb.InfoServer
}

// Server struct implements the EmmyServer interface.
var _ EmmyServer = (*Server)(nil)

type Server struct {
	GrpcServer *grpc.Server
	Logger     log.Logger
	SessionManager
	RegistrationManager
	clRecordManager cl.ReceiverRecordManager
}

// NewServer initializes an instance of the Server struct and returns a pointer.
// It performs some default configuration (tracing of gRPC communication and interceptors)
// and registers RPC server handlers with gRPC server. It requires TLS cert and keyfile
// in order to establish a secure channel with clients.
func NewServer(certFile, keyFile string, regMgr RegistrationManager,
	recMgr cl.ReceiverRecordManager, logger log.Logger) (*Server, error) {
	logger.Info("Instantiating new server")

	// Obtain TLS credentials
	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	logger.Infof("Successfully read certificate [%s] and key [%s]", certFile, keyFile)

	sessionManager, err := NewRandSessionKeyGen(config.LoadSessionKeyMinByteLen())
	if err != nil {
		logger.Warning(err)
	}

	// Allow as much concurrent streams as possible and register a gRPC stream interceptor
	// for logging and monitoring purposes.
	server := &Server{
		GrpcServer: grpc.NewServer(
			grpc.Creds(creds),
			grpc.MaxConcurrentStreams(math.MaxUint32),
			grpc.StreamInterceptor(grpc_prometheus.StreamServerInterceptor),
		),
		Logger:              logger,
		SessionManager:      sessionManager,
		RegistrationManager: regMgr,
		clRecordManager:     recMgr,
	}

	// Disable tracing by default, as is used for debugging purposes.
	// The user will be able to turn it on via Server's EnableTracing function.
	grpc.EnableTracing = false

	// Register our services with the supporting gRPC server
	server.registerServices()

	// Initialize gRPC metrics offered by Prometheus package
	grpc_prometheus.Register(server.GrpcServer)

	return server, nil
}

// Start configures and starts the protocol server at the requested port.
func (s *Server) Start(port int) error {
	connStr := fmt.Sprintf(":%d", port)
	listener, err := net.Listen("tcp", connStr)
	if err != nil {
		return fmt.Errorf("could not connect: %v", err)
	}

	// Register Prometheus metrics handler and serve metrics page on the desired endpoint.
	// Metrics are handled via HTTP in a separate goroutine as gRPC requests,
	// as grpc server's performance over HTTP (GrpcServer.ServeHTTP) is much worse.
	http.Handle("/metrics", promhttp.Handler())

	// After this, /metrics will be available, along with /debug/requests, /debug/events in
	// case server's EnableTracing function is called.
	go http.ListenAndServe(":8881", nil)

	// From here on, gRPC server will accept connections
	s.Logger.Noticef("emmy server listening for connections on port %d", port)
	s.GrpcServer.Serve(listener)
	return nil
}

// Teardown stops the protocol server by gracefully stopping enclosed gRPC server.
func (s *Server) Teardown() {
	s.Logger.Notice("Tearing down gRPC server")
	s.GrpcServer.GracefulStop()
}

// EnableTracing instructs the gRPC framework to enable its tracing capability, which
// is mainly used for debugging purposes.
// Although this function does not explicitly affect the Server struct, it is wired to Server
// in order to provide a nicer API when setting up the server.
func (s *Server) EnableTracing() {
	grpc.EnableTracing = true
	s.Logger.Notice("Enabled gRPC tracing")
}

// registerServices binds gRPC server interfaces to the server instance itself, as the server
// provides implementations of these interfaces.
func (s *Server) registerServices() {
	pb.RegisterInfoServer(s.GrpcServer, s)
	pb.RegisterPseudonymSystemServer(s.GrpcServer, s)
	pb.RegisterPseudonymSystemCAServer(s.GrpcServer, s)
	pb.RegisterCLServer(s.GrpcServer, s)

	s.Logger.Notice("Registered gRPC Services")
}

func (s *Server) send(msg *pb.Message, stream zk.ServerStream) error {
	if err := stream.Send(msg); err != nil {
		return fmt.Errorf("error sending message: %v", err)
	}

	s.Logger.Infof("Successfully sent response of type %T", msg.Content)
	s.Logger.Debugf("%+v", msg)

	return nil
}

func (s *Server) receive(stream zk.ServerStream) (*pb.Message, error) {
	resp, err := stream.Recv()
	if err == io.EOF {
		return nil, err
	} else if err != nil {
		return nil, fmt.Errorf("an error occurred: %v", err)
	}
	s.Logger.Infof("Received request of type %T from the stream", resp.Content)
	s.Logger.Debugf("%+v", resp)

	return resp, nil
}

func (s *Server) GetServiceInfo(ctx context.Context, _ *empty.Empty) (*pb.ServiceInfo, error) {
	s.Logger.Info("Client requested service information")

	name, provider, description := config.LoadServiceInfo()
	info := &pb.ServiceInfo{
		Name:        name,
		Provider:    provider,
		Description: description,
	}

	return info, nil
}

// RegistrationManager checks for the presence of a registration key,
// removing it in case it exists.
// The bolean return argument indicates success (registration key
// present and subsequently deleted) or failure (absence of registration
// key).
type RegistrationManager interface {
	CheckRegistrationKey(string) (bool, error)
}

type RedisClient struct {
	*redis.Client
}

func NewRedisClient(c *redis.Client) *RedisClient {
	return &RedisClient{
		Client: c,
	}
}

// CheckRegistrationKey checks whether provided key is present in registration database and deletes it,
// preventing another registration with the same key.
// Returns true if key was present (registration allowed), false otherwise.
func (c *RedisClient) CheckRegistrationKey(key string) (bool, error) {
	resp := c.Del(key)

	err := resp.Err()

	if err != nil {
		return false, err
	}

	return resp.Val() == 1, nil // one deleted entry indicates that the key was present in the DB
}

// SessionManager generates a new session key.
// It returns a string containing the generated session key
// or an error in case session key could not be generated.
type SessionManager interface {
	GenerateSessionKey() (*string, error)
}

// MIN_SESSION_KEY_BYTE_LEN represents the minimal allowed length
// of the session key in bytes, for security reasons.
const MIN_SESSION_KEY_BYTE_LEN = 24

// RandSessionKeyGen generates session keys of the desired byte
// length from random bytes.
type RandSessionKeyGen struct {
	byteLen int
}

// NewRandSessionKeyGen creates a new RandSessionKeyGen instance.
// The new instance will be configured to generate session keys
// with exactly byteLen bytes. For security reasons, the function
// checks the byteLen against the value of MIN_SESSION_KEY_BYTE_LEN.
// If the provided byteLen is smaller than MIN_SESSION_KEY_BYTE_LEN,
// an error is set and the returned RandSessionKeyGen is configured
// to use MIN_SESSION_KEY_BYTE_LEN instead of the provided byteLen.
func NewRandSessionKeyGen(byteLen int) (*RandSessionKeyGen, error) {
	var err error
	if byteLen < MIN_SESSION_KEY_BYTE_LEN {
		err = fmt.Errorf("desired length of the session key (%d B) is too short, falling back to %d B",
			byteLen, MIN_SESSION_KEY_BYTE_LEN)
		byteLen = MIN_SESSION_KEY_BYTE_LEN
	}
	return &RandSessionKeyGen{
		byteLen: byteLen,
	}, err
}

// GenerateSessionKey produces a secure random session key and returns
// its base64-encoded representation that is URL-safe.
// It reports an error in case random byte sequence could not be generated.
func (m *RandSessionKeyGen) GenerateSessionKey() (*string, error) {
	randBytes := make([]byte, m.byteLen)

	// reads m.byteLen random bytes (e.g. len(randBytes)) to randBytes array
	_, err := rand.Read(randBytes)

	// an error may occur if the system's secure RNG doesn't function properly, in which case
	// we can't generate a secure session key
	if err != nil {
		return nil, err
	}

	sessionKey := base64.URLEncoding.EncodeToString(randBytes)
	return &sessionKey, nil
}

func (s *Server) GenerateNym(stream pb.PseudonymSystem_GenerateNymServer) error {
	req, err := s.receive(stream)
	if err != nil {
		return err
	}

	group := config.LoadSchnorrGroup()
	caPubKey := config.LoadPseudonymsysCAPubKey()
	org := pseudsys.NewNymGenerator(group, caPubKey)

	proofRandData := req.GetPseudonymsysNymGenProofRandomData()
	x1 := new(big.Int).SetBytes(proofRandData.X1)
	nymA := new(big.Int).SetBytes(proofRandData.A1)
	nymB := new(big.Int).SetBytes(proofRandData.B1)
	x2 := new(big.Int).SetBytes(proofRandData.X2)
	blindedA := new(big.Int).SetBytes(proofRandData.A2)
	blindedB := new(big.Int).SetBytes(proofRandData.B2)
	signatureR := new(big.Int).SetBytes(proofRandData.R)
	signatureS := new(big.Int).SetBytes(proofRandData.S)

	regKeyOk, err := s.RegistrationManager.CheckRegistrationKey(proofRandData.RegKey)

	var resp *pb.Message

	if !regKeyOk || err != nil {
		s.Logger.Debugf("registration key %s ok=%t, error=%v",
			proofRandData.RegKey, regKeyOk, err)
		return status.Error(codes.NotFound, "registration key verification failed")
	}

	challenge, err := org.GetChallenge(nymA, blindedA, nymB, blindedB, x1, x2, signatureR, signatureS)
	if err != nil {
		s.Logger.Debug(err)
		return status.Error(codes.Internal, err.Error())

	}
	resp = &pb.Message{
		Content: &pb.Message_PedersenDecommitment{
			&pb.PedersenDecommitment{
				X: challenge.Bytes(),
			},
		},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	req, err = s.receive(stream)
	if err != nil {
		return err
	}

	proofData := req.GetSchnorrProofData() // SchnorrProofData is used in DLog equality proof as well
	z := new(big.Int).SetBytes(proofData.Z)
	valid := org.Verify(z)

	resp = &pb.Message{
		Content: &pb.Message_Status{&pb.Status{Success: valid}},
	}

	if err = s.send(resp, stream); err != nil {
		return err
	}

	return nil
}

func (s *Server) ObtainCredential(stream pb.PseudonymSystem_ObtainCredentialServer) error {
	req, err := s.receive(stream)
	if err != nil {
		return err
	}

	group := config.LoadSchnorrGroup()
	secKey := config.LoadPseudonymsysOrgSecrets("org1", "dlog")
	org := pseudsys.NewCredIssuer(group, secKey)

	sProofRandData := req.GetSchnorrProofRandomData()
	x := new(big.Int).SetBytes(sProofRandData.X)
	a := new(big.Int).SetBytes(sProofRandData.A)
	b := new(big.Int).SetBytes(sProofRandData.B)
	challenge := org.GetChallenge(a, b, x)

	resp := &pb.Message{
		Content: &pb.Message_Bigint{
			&pb.BigInt{
				X1: challenge.Bytes(),
			},
		},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	req, err = s.receive(stream)
	if err != nil {
		return err
	}

	proofData := req.GetBigint()
	z := new(big.Int).SetBytes(proofData.X1)

	x11, x12, x21, x22, A, B, err := org.Verify(z)
	if err != nil {
		s.Logger.Debug(err)
		return status.Error(codes.Internal, err.Error())
	}
	resp = &pb.Message{
		Content: &pb.Message_PseudonymsysIssueProofRandomData{
			&pb.PseudonymsysIssueProofRandomData{
				X11: x11.Bytes(),
				X12: x12.Bytes(),
				X21: x21.Bytes(),
				X22: x22.Bytes(),
				A:   A.Bytes(),
				B:   B.Bytes(),
			},
		},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	req, err = s.receive(stream)
	if err != nil {
		return err
	}

	challenges := req.GetDoubleBigint()
	challenge1 := new(big.Int).SetBytes(challenges.X1)
	challenge2 := new(big.Int).SetBytes(challenges.X2)

	z1, z2 := org.GetProofData(challenge1, challenge2)
	resp = &pb.Message{
		Content: &pb.Message_DoubleBigint{
			&pb.DoubleBigInt{
				X1: z1.Bytes(),
				X2: z2.Bytes(),
			},
		},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	return nil
}

func (s *Server) TransferCredential(stream pb.PseudonymSystem_TransferCredentialServer) error {
	req, err := s.receive(stream)
	if err != nil {
		return err
	}

	group := config.LoadSchnorrGroup()
	secKey := config.LoadPseudonymsysOrgSecrets("org1", "dlog")
	org := pseudsys.NewCredVerifier(group, secKey)

	data := req.GetPseudonymsysTransferCredentialData()
	orgName := data.OrgName
	x1 := new(big.Int).SetBytes(data.X1)
	x2 := new(big.Int).SetBytes(data.X2)
	nymA := new(big.Int).SetBytes(data.NymA)
	nymB := new(big.Int).SetBytes(data.NymB)

	t1 := schnorr.NewBlindedTrans(
		new(big.Int).SetBytes(data.Credential.T1.A),
		new(big.Int).SetBytes(data.Credential.T1.B),
		new(big.Int).SetBytes(data.Credential.T1.Hash),
		new(big.Int).SetBytes(data.Credential.T1.ZAlpha),
	)

	t2 := schnorr.NewBlindedTrans(
		new(big.Int).SetBytes(data.Credential.T2.A),
		new(big.Int).SetBytes(data.Credential.T2.B),
		new(big.Int).SetBytes(data.Credential.T2.Hash),
		new(big.Int).SetBytes(data.Credential.T2.ZAlpha),
	)

	credential := pseudsys.NewCred(
		new(big.Int).SetBytes(data.Credential.SmallAToGamma),
		new(big.Int).SetBytes(data.Credential.SmallBToGamma),
		new(big.Int).SetBytes(data.Credential.AToGamma),
		new(big.Int).SetBytes(data.Credential.BToGamma),
		t1, t2,
	)

	challenge := org.GetChallenge(nymA, nymB,
		credential.SmallAToGamma, credential.SmallBToGamma, x1, x2)

	resp := &pb.Message{
		Content: &pb.Message_Bigint{
			&pb.BigInt{
				X1: challenge.Bytes(),
			},
		},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	req, err = s.receive(stream)
	if err != nil {
		return err
	}

	// PubKeys of the organization that issue a credential:
	orgPubKeys := config.LoadPseudonymsysOrgPubKeys(orgName)

	proofData := req.GetBigint()
	z := new(big.Int).SetBytes(proofData.X1)

	if verified := org.Verify(z, credential, orgPubKeys); !verified {
		s.Logger.Debug("User authentication failed")
		return status.Error(codes.Unauthenticated, "user authentication failed")
	}

	sessionKey, err := s.GenerateSessionKey()
	if err != nil {
		s.Logger.Debug(err)
		return status.Error(codes.Internal, "failed to obtain session key")
	}

	resp = &pb.Message{
		Content: &pb.Message_SessionKey{
			SessionKey: &pb.SessionKey{
				Value: *sessionKey,
			},
		},
	}

	if err = s.send(resp, stream); err != nil {
		return err
	}

	return nil
}

func (s *Server) GenerateNym_EC(stream pb.PseudonymSystem_GenerateNym_ECServer) error {
	req, err := s.receive(stream)
	if err != nil {
		return err
	}

	caPubKey := config.LoadPseudonymsysCAPubKey()
	org := ecpseudsys.NewNymGenerator(caPubKey, curve)

	proofRandData := req.GetPseudonymsysNymGenProofRandomDataEc()
	x1 := zk.GetNativeType(proofRandData.X1)
	nymA := zk.GetNativeType(proofRandData.A1)
	nymB := zk.GetNativeType(proofRandData.B1)
	x2 := zk.GetNativeType(proofRandData.X2)
	blindedA := zk.GetNativeType(proofRandData.A2)
	blindedB := zk.GetNativeType(proofRandData.B2)
	signatureR := new(big.Int).SetBytes(proofRandData.R)
	signatureS := new(big.Int).SetBytes(proofRandData.S)

	regKeyOk, err := s.RegistrationManager.CheckRegistrationKey(proofRandData.RegKey)

	var resp *pb.Message

	if !regKeyOk || err != nil {
		s.Logger.Debugf("Registration key %s ok=%t, error=%v",
			proofRandData.RegKey, regKeyOk, err)
		return status.Error(codes.NotFound, "registration key verification failed")

	}
	challenge, err := org.GetChallenge(nymA, blindedA, nymB, blindedB, x1, x2, signatureR, signatureS)
	if err != nil {
		s.Logger.Debug(err)
		return status.Error(codes.Internal, err.Error())
	}
	resp = &pb.Message{
		Content: &pb.Message_PedersenDecommitment{
			&pb.PedersenDecommitment{
				X: challenge.Bytes(),
			},
		},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	req, err = s.receive(stream)
	if err != nil {
		return err
	}

	proofData := req.GetSchnorrProofData() // SchnorrProofData is used in DLog equality proof as well
	z := new(big.Int).SetBytes(proofData.Z)
	valid := org.Verify(z)

	resp = &pb.Message{
		Content: &pb.Message_Status{&pb.Status{Success: valid}},
	}

	if err = s.send(resp, stream); err != nil {
		return err
	}

	return nil
}

func (s *Server) ObtainCredential_EC(stream pb.PseudonymSystem_ObtainCredential_ECServer) error {
	req, err := s.receive(stream)
	if err != nil {
		return err
	}

	proofRandData := req.GetSchnorrEcProofRandomData()
	x := zk.GetNativeType(proofRandData.X)
	a := zk.GetNativeType(proofRandData.A)
	b := zk.GetNativeType(proofRandData.B)

	secKey := config.LoadPseudonymsysOrgSecrets("org1", "ecdlog")
	org := ecpseudsys.NewCredIssuer(secKey, curve)
	challenge := org.GetChallenge(a, b, x)

	resp := &pb.Message{
		Content: &pb.Message_Bigint{
			&pb.BigInt{
				X1: challenge.Bytes(),
			},
		},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	req, err = s.receive(stream)
	if err != nil {
		return err
	}

	proofData := req.GetBigint()
	z := new(big.Int).SetBytes(proofData.X1)

	x11, x12, x21, x22, A, B, err := org.Verify(z)

	if err != nil {
		s.Logger.Debug(err)
		return status.Error(codes.Internal, err.Error())
	}
	resp = &pb.Message{
		Content: &pb.Message_PseudonymsysIssueProofRandomDataEc{
			&pb.PseudonymsysIssueProofRandomDataEC{
				X11: zk.ToPbECGroupElement(x11),
				X12: zk.ToPbECGroupElement(x12),
				X21: zk.ToPbECGroupElement(x21),
				X22: zk.ToPbECGroupElement(x22),
				A:   zk.ToPbECGroupElement(A),
				B:   zk.ToPbECGroupElement(B),
			},
		},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	req, err = s.receive(stream)
	if err != nil {
		return err
	}

	challenges := req.GetDoubleBigint()
	challenge1 := new(big.Int).SetBytes(challenges.X1)
	challenge2 := new(big.Int).SetBytes(challenges.X2)

	z1, z2 := org.GetProofData(challenge1, challenge2)
	resp = &pb.Message{
		Content: &pb.Message_DoubleBigint{
			&pb.DoubleBigInt{
				X1: z1.Bytes(),
				X2: z2.Bytes(),
			},
		},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	return nil
}

func (s *Server) TransferCredential_EC(stream pb.PseudonymSystem_TransferCredential_ECServer) error {
	req, err := s.receive(stream)
	if err != nil {
		return err
	}

	secKey := config.LoadPseudonymsysOrgSecrets("org1", "ecdlog")
	org := ecpseudsys.NewCredVerifier(secKey, curve)

	data := req.GetPseudonymsysTransferCredentialDataEc()
	orgName := data.OrgName
	x1 := zk.GetNativeType(data.X1)
	x2 := zk.GetNativeType(data.X2)
	nymA := zk.GetNativeType(data.NymA)
	nymB := zk.GetNativeType(data.NymB)

	t1 := ecschnorr.NewBlindedTrans(
		new(big.Int).SetBytes(data.Credential.T1.A.X),
		new(big.Int).SetBytes(data.Credential.T1.A.Y),
		new(big.Int).SetBytes(data.Credential.T1.B.X),
		new(big.Int).SetBytes(data.Credential.T1.B.Y),
		new(big.Int).SetBytes(data.Credential.T1.Hash),
		new(big.Int).SetBytes(data.Credential.T1.ZAlpha))

	t2 := ecschnorr.NewBlindedTrans(
		new(big.Int).SetBytes(data.Credential.T2.A.X),
		new(big.Int).SetBytes(data.Credential.T2.A.Y),
		new(big.Int).SetBytes(data.Credential.T2.B.X),
		new(big.Int).SetBytes(data.Credential.T2.B.Y),
		new(big.Int).SetBytes(data.Credential.T2.Hash),
		new(big.Int).SetBytes(data.Credential.T2.ZAlpha))

	credential := ecpseudsys.NewCred(
		zk.GetNativeType(data.Credential.SmallAToGamma),
		zk.GetNativeType(data.Credential.SmallBToGamma),
		zk.GetNativeType(data.Credential.AToGamma),
		zk.GetNativeType(data.Credential.BToGamma),
		t1, t2,
	)

	challenge := org.GetChallenge(nymA, nymB,
		credential.SmallAToGamma, credential.SmallBToGamma, x1, x2)

	resp := &pb.Message{
		Content: &pb.Message_Bigint{
			&pb.BigInt{
				X1: challenge.Bytes(),
			},
		},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	req, err = s.receive(stream)
	if err != nil {
		return err
	}

	// PubKeys of the organization that issue a credential:
	orgPubKeys := config.LoadPseudonymsysOrgPubKeysEC(orgName)

	proofData := req.GetBigint()
	z := new(big.Int).SetBytes(proofData.X1)

	if verified := org.Verify(z, credential, orgPubKeys); !verified {
		s.Logger.Debug("User authentication failed")
		return status.Error(codes.Unauthenticated, "user authentication failed")
	}

	sessionKey, err := s.GenerateSessionKey()
	if err != nil {
		s.Logger.Debug(err)
		return status.Error(codes.Internal, "failed to obtain session key")
	}

	resp = &pb.Message{
		Content: &pb.Message_SessionKey{
			SessionKey: &pb.SessionKey{
				Value: *sessionKey,
			},
		},
	}

	if err = s.send(resp, stream); err != nil {
		return err
	}

	return nil
}

func (s *Server) GenerateCertificate(stream pb.PseudonymSystemCA_GenerateCertificateServer) error {
	var err error

	req, err := s.receive(stream)
	if err != nil {
		return err
	}

	group := config.LoadSchnorrGroup()
	d := config.LoadPseudonymsysCASecret()
	pubKey := config.LoadPseudonymsysCAPubKey()
	ca := pseudsys.NewCA(group, d, pubKey)

	sProofRandData := req.GetSchnorrProofRandomData()
	x := new(big.Int).SetBytes(sProofRandData.X)
	a := new(big.Int).SetBytes(sProofRandData.A)
	b := new(big.Int).SetBytes(sProofRandData.B)

	challenge := ca.GetChallenge(a, b, x)
	resp := &pb.Message{
		Content: &pb.Message_Bigint{
			&pb.BigInt{
				X1: challenge.Bytes(),
			},
		},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	req, err = s.receive(stream)
	if err != nil {
		return err
	}

	sProofData := req.GetSchnorrProofData()
	z := new(big.Int).SetBytes(sProofData.Z)
	cert, err := ca.Verify(z)

	if err != nil {
		s.Logger.Debug(err)
		return status.Error(codes.Internal, err.Error())
	}

	resp = &pb.Message{
		Content: &pb.Message_PseudonymsysCaCertificate{
			&pb.PseudonymsysCACertificate{
				BlindedA: cert.BlindedA.Bytes(),
				BlindedB: cert.BlindedB.Bytes(),
				R:        cert.R.Bytes(),
				S:        cert.S.Bytes(),
			},
		},
	}

	if err = s.send(resp, stream); err != nil {
		return err
	}

	return nil
}

func (s *Server) GenerateCertificate_EC(stream pb.PseudonymSystemCA_GenerateCertificate_ECServer) error {
	req, err := s.receive(stream)
	if err != nil {
		return err
	}

	d := config.LoadPseudonymsysCASecret()
	pubKey := config.LoadPseudonymsysCAPubKey()
	ca := ecpseudsys.NewCA(d, pubKey, curve)

	sProofRandData := req.GetSchnorrEcProofRandomData()
	x := zk.GetNativeType(sProofRandData.X)
	a := zk.GetNativeType(sProofRandData.A)
	b := zk.GetNativeType(sProofRandData.B)

	challenge := ca.GetChallenge(a, b, x)
	resp := &pb.Message{
		Content: &pb.Message_Bigint{
			&pb.BigInt{
				X1: challenge.Bytes(),
			},
		},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	req, err = s.receive(stream)
	if err != nil {
		return err
	}

	sProofData := req.GetSchnorrProofData()
	z := new(big.Int).SetBytes(sProofData.Z)
	cert, err := ca.Verify(z)

	if err != nil {
		s.Logger.Debug(err)
		return status.Error(codes.Internal, err.Error())
	}

	resp = &pb.Message{
		Content: &pb.Message_PseudonymsysCaCertificateEc{
			&pb.PseudonymsysCACertificateEC{
				BlindedA: zk.ToPbECGroupElement(cert.BlindedA),
				BlindedB: zk.ToPbECGroupElement(cert.BlindedB),
				R:        cert.R.Bytes(),
				S:        cert.S.Bytes(),
			},
		},
	}

	if err = s.send(resp, stream); err != nil {
		return err
	}

	return nil
}

func (s *Server) GetCredentialStructure(ctx context.Context, _ *empty.Empty) (*pb.CredStructure, error) {
	s.Logger.Info("Client requested credential structure information")

	structure, err := config.LoadCredentialStructure()
	if err != nil {
		return nil, err
	}

	attrs, attrCount, err := cl.ParseAttrs(structure)
	if err != nil {
		return nil, err
	}
	credAttrs := make([]*pb.CredAttribute, len(attrs))

	for i, a := range attrs {
		attr := &pb.Attribute{
			Name:  a.GetName(),
			Known: a.IsKnown(),
		}
		switch a.(type) {
		case *cl.StrAttr:
			credAttrs[i] = &pb.CredAttribute{
				Type: &pb.CredAttribute_StringAttr{
					StringAttr: &pb.StringAttribute{
						Attr: attr,
					},
				},
			}
		case *cl.Int64Attr:
			credAttrs[i] = &pb.CredAttribute{
				Type: &pb.CredAttribute_IntAttr{
					IntAttr: &pb.IntAttribute{
						Attr: attr,
					},
				},
			}
		}
	}

	return &pb.CredStructure{
		NKnown:     int32(attrCount.Known),
		NCommitted: int32(attrCount.Committed),
		NHidden:    int32(attrCount.Hidden),
		Attributes: credAttrs,
	}, nil
}

func (s *Server) GetAcceptableCredentials(ctx context.Context, _ *empty.Empty) (*pb.AcceptableCreds, error) {
	s.Logger.Info("Client requested acceptable credentials information")
	accCreds, err := config.LoadAcceptableCredentials()
	if err != nil {
		return nil, err
	}

	var credentials []*pb.AcceptableCred
	for name, attrs := range accCreds {
		cred := &pb.AcceptableCred{
			OrgName:       name,
			RevealedAttrs: attrs,
		}
		credentials = append(credentials, cred)
	}

	return &pb.AcceptableCreds{
		Creds: credentials,
	}, nil
}

func (s *Server) IssueCredential(stream pb.CL_IssueCredentialServer) error {
	req, err := s.receive(stream)
	if err != nil {
		return err
	}

	// TOD0: for known attributes IssueCredential should fill the values - attributes
	// are stored (under registration key) in the DB and then obtained by Org.

	initReq := req.GetRegKey()
	regKeyOk, err := s.RegistrationManager.CheckRegistrationKey(initReq.RegKey)
	if !regKeyOk || err != nil {
		s.Logger.Debugf("registration key %s ok=%t, error=%v",
			initReq.RegKey, regKeyOk, err)
		return status.Error(codes.NotFound, "registration key verification failed")
	}

	org, err := cl.LoadOrg("../client/testdata/clPubKey.gob", "../client/testdata/clSecKey.gob")
	if err != nil {
		return err
	}

	nonce := org.GetCredIssueNonce()
	resp := &pb.Message{
		Content: &pb.Message_Bigint{
			&pb.BigInt{
				X1: nonce.Bytes(),
			},
		},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	req, err = s.receive(stream)
	if err != nil {
		return err
	}

	cReq := req.GetCLCredReq()
	credReq, err := zk.GetNativeTypeFromCredReq(cReq)
	if err != nil {
		return err
	}

	// Issue the credential
	res, err := org.IssueCred(credReq)
	if err != nil {
		return fmt.Errorf("error when issuing credential: %v", err)
	}
	// Store the newly obtained receiver record to the database
	if err = s.clRecordManager.Store(credReq.Nym, res.Record); err != nil {
		return err
	}

	pbCred := zk.ToPbCLCredential(res.Cred, res.AProof)
	resp = &pb.Message{
		Content: &pb.Message_CLCredential{pbCred},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	return nil
}

func (s *Server) UpdateCredential(stream pb.CL_UpdateCredentialServer) error {
	req, err := s.receive(stream)
	if err != nil {
		return err
	}

	org, err := cl.LoadOrg("../client/testdata/clPubKey.gob", "../client/testdata/clSecKey.gob")
	if err != nil {
		return err
	}

	u := req.GetUpdateClCredential()
	nym, nonce, newKnownAttrs := zk.GetNativeTypeFromUpdateCredential(u)

	// Retrieve the receiver record from the database
	rec, err := s.clRecordManager.Load(nym)
	if err != nil {
		return err
	}
	// Do credential update
	res, err := org.UpdateCred(nym, rec, nonce, newKnownAttrs)
	if err != nil {
		return fmt.Errorf("error when updating credential: %v", err)
	}
	// Store the updated receiver record to the database
	if err = s.clRecordManager.Store(nym, res.Record); err != nil {
		return err
	}

	pbCred := zk.ToPbCLCredential(res.Cred, res.AProof)
	resp := &pb.Message{
		Content: &pb.Message_CLCredential{pbCred},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	return nil
}

func (s *Server) ProveCredential(stream pb.CL_ProveCredentialServer) error {
	req, err := s.receive(stream)
	if err != nil {
		return err
	}

	org, err := cl.LoadOrg("../client/testdata/clPubKey.gob", "../client/testdata/clSecKey.gob")
	if err != nil {
		return err
	}

	nonce := org.GetProveCredNonce()
	resp := &pb.Message{
		Content: &pb.Message_Bigint{
			&pb.BigInt{
				X1: nonce.Bytes(),
			},
		},
	}

	if err := s.send(resp, stream); err != nil {
		return err
	}

	req, err = s.receive(stream)
	if err != nil {
		return err
	}

	pReq := req.GetProveClCredential()
	A, proof, knownAttrs, commitmentsOfAttrs, revealedKnownAttrsIndices,
		revealedCommitmentsOfAttrsIndices, err := zk.GetNativeTypeFromProveCredential(pReq)
	if err != nil {
		return err
	}

	verified, err := org.ProveCred(A, proof, revealedKnownAttrsIndices,
		revealedCommitmentsOfAttrsIndices, knownAttrs, commitmentsOfAttrs)
	if err != nil {
		s.Logger.Debug(err)
		return status.Error(codes.Internal, "error when proving credential")
	}

	if !verified {
		s.Logger.Debug("User authentication failed")
		return status.Error(codes.Unauthenticated, "user authentication failed")
	}

	sessionKey, err := s.GenerateSessionKey()
	if err != nil {
		s.Logger.Debug(err)
		return status.Error(codes.Internal, "failed to obtain session key")
	}

	// TODO: here session key needs to be stored to enable validation

	resp = &pb.Message{
		Content: &pb.Message_SessionKey{
			SessionKey: &pb.SessionKey{
				Value: *sessionKey,
			},
		},
	}

	if err = s.send(resp, stream); err != nil {
		return err
	}

	return nil
}
