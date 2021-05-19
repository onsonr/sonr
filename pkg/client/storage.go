package client

import (
	"bytes"
	"context"
	"io"

	md "github.com/sonr-io/core/pkg/models"
	"google.golang.org/protobuf/proto"
	"storj.io/uplink"
)

type Storage struct {
	ctx       context.Context
	appAccess *uplink.Access
}

// @ Start New Storage Uplink
func NewUplink(ctx context.Context, appAPIKey string, rootPassword string) (*Storage, error) {
	appAccess, err := uplink.ParseAccess(appAPIKey)
	if err != nil {
		return nil, err
	}

	return &Storage{
		ctx:       ctx,
		appAccess: appAccess,
	}, nil
}

// @ Get User from Remote Data Store
func (s *Storage) GetUser(id string) (*md.User, error) {
	project, err := uplink.OpenProject(s.ctx, s.appAccess)
	if err != nil {
		return nil, err
	}
	defer project.Close()

	buffer := new(bytes.Buffer)
	d, err := project.DownloadObject(s.ctx, "users", id, &uplink.DownloadOptions{})
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(buffer, d)
	if err != nil {
		return nil, err
	}

	user := &md.User{}
	err = proto.Unmarshal(buffer.Bytes(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// @ Put User in Remote Data Store
func (s *Storage) PutUser(u *md.User) error {
	// Open PRoject
	project, err := uplink.OpenProject(s.ctx, s.appAccess)
	if err != nil {
		return err
	}
	defer project.Close()

	// Marshal Userl
	bytes, err := proto.Marshal(u)
	if err != nil {
		return err
	}

	// Create Upload Object
	object, err := project.UploadObject(s.ctx, "users", u.GetId(), &uplink.UploadOptions{})
	if err != nil {
		return err
	}

	// Write User Bytes to Object
	_, err = object.Write(bytes)
	if err != nil {
		return err
	}

	// Push Changes
	err = object.Commit()
	if err != nil {
		return err
	}
	return nil
}
