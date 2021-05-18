package host

import (
	"context"
	"errors"

	"github.com/libp2p/go-libp2p-core/crypto"
	md "github.com/sonr-io/core/pkg/models"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/go-threads/db"
	"google.golang.org/grpc"
)

// @ Initializes Database for this Host
func (hn *HostNode) InitDB(privateKey crypto.PrivKey) error {
	client, err := client.NewClient("/ip4/127.0.0.1/tcp/6006", grpc.WithInsecure())
	if err != nil {
		return err
	}
	hn.DBClient = client

	hn.DBThreadID = thread.NewIDV1(thread.Raw, 32)
	threadToken, err := client.GetToken(context.Background(), thread.NewLibp2pIdentity(privateKey))
	if err != nil {
		return err
	}
	hn.DBToken = threadToken

	threadAddr, err := GetRemoteThreadMultiAddr()
	if err != nil {
		return err
	}

	err = hn.DBClient.NewDBFromAddr(hn.ctx, threadAddr, thread.NewRandomKey())
	if err != nil {
		return err
	}

	config := db.CollectionConfig{
		Name:   "Users",
		Schema: md.ReflectUserToJson(),
		Indexes: []db.Index{{
			Path:   "id",
			Unique: true,
		}},
	}

	err = hn.DBClient.NewCollection(hn.ctx, hn.DBThreadID, config)
	if err != nil {
		return err
	}
	return nil
}

// @ Read User from Database
func (hn *HostNode) ReadUser(id string) (*md.User, error) {
	// Determine if an instance exists by ID
	exists, err := hn.DBClient.Has(context.Background(), hn.DBThreadID, "Users", []string{id})
	if err != nil {
		return nil, err
	}

	// Check if Value Exists
	if exists {
		user := &md.User{}
		err = hn.DBClient.FindByID(context.Background(), hn.DBThreadID, "Users", id, user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, errors.New("User does not Exist.")
}

// @ Writes User to Database
func (hn *HostNode) WriteUser(u *md.User) error {
	// Determine if an instance exists by ID
	exists, err := hn.DBClient.Has(context.Background(), hn.DBThreadID, "Users", []string{u.GetId()})
	if err != nil {
		return err
	}

	// Update Value if User Exists
	if exists {
		err := hn.DBClient.Save(context.Background(), hn.DBThreadID, "Users", client.Instances{u})
		if err != nil {
			return err
		}
	} else {
		// Create Value if User Doesnt Exist
		ids, err := hn.DBClient.Create(context.Background(), hn.DBThreadID, "Users", client.Instances{u})
		if err != nil {
			return err
		}
		hn.IDS = ids
	}
	return nil
}
