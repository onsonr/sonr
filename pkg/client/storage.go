package client

import (
	"bytes"
	"context"
	"fmt"
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
	// Create User Access
	userAccess, err := s.appAccess.Share(uplink.FullPermission())
	if err != nil {
		return nil, err
	}

	// Open Project
	project, err := uplink.OpenProject(s.ctx, userAccess)
	if err != nil {
		return nil, err
	}
	defer project.Close()

	// Marshal USer
	buffer := new(bytes.Buffer)
	d, err := project.DownloadObject(s.ctx, "users", id, &uplink.DownloadOptions{})
	if err != nil {
		return nil, err
	}

	// Copy to Buffer
	_, err = io.Copy(buffer, d)
	if err != nil {
		return nil, err
	}

	// Convert to Object
	user := &md.User{}
	err = proto.Unmarshal(buffer.Bytes(), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// @ Put User in Remote Data Store
func (s *Storage) PutUser(u *md.User) error {
	// Open Project
	project, err := uplink.OpenProject(s.ctx, s.appAccess)
	if err != nil {
		return fmt.Errorf("could not open project: %v", err)
	}
	defer project.Close()

	// Ensure the desired Bucket within the Project is created.
	_, err = project.EnsureBucket(s.ctx, "users")
	if err != nil {
		return fmt.Errorf("could not ensure bucket: %v", err)
	}

	// Create Upload Object
	object, err := project.UploadObject(s.ctx, "users", u.GetId(), nil)
	if err != nil {
		return err
	}

	// Marshal User
	data, err := proto.Marshal(u)
	if err != nil {
		return err
	}

	// Copy to Object
	buf := bytes.NewBuffer(data)
	_, err = io.Copy(object, buf)
	if err != nil {
		_ = object.Abort()
		return fmt.Errorf("could not upload data: %v", err)
	}

	// Push Changes
	err = object.Commit()
	if err != nil {
		return err
	}
	return nil
}
