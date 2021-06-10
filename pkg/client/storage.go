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

// @ Get User from Remote Data Store
func GetUser(ctx context.Context, appAPIKey string, id string) (*md.User, error) {
	// Parse Access
	appAccess, err := uplink.ParseAccess(appAPIKey)
	if err != nil {
		return nil, err
	}

	// Open Project
	project, err := uplink.OpenProject(ctx, appAccess)
	if err != nil {
		return nil, err
	}
	defer project.Close()

	// Create Download Object
	d, err := project.DownloadObject(ctx, "users", id, nil)
	if err != nil {
		return nil, err
	}

	// Copy to Buffer
	buffer := new(bytes.Buffer)
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
func PutUser(ctx context.Context, appAPIKey string, u *md.User) error {
	// Parse Access
	appAccess, err := uplink.ParseAccess(appAPIKey)
	if err != nil {
		return err
	}

	// Open Project
	project, err := uplink.OpenProject(ctx, appAccess)
	if err != nil {
		return fmt.Errorf("could not open project: %v", err)
	}
	defer project.Close()

	// Ensure the desired Bucket within the Project is created.
	_, err = project.EnsureBucket(ctx, "users")
	if err != nil {
		return fmt.Errorf("could not ensure bucket: %v", err)
	}

	// Get Prefix
	prefix, err := u.Prefix()
	if err != nil {
		return err
	}

	// Create Upload Object
	object, err := project.UploadObject(ctx, "users", prefix, nil)
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
