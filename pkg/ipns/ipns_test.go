package ipns_test

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sonr-io/sonr/internal/bucket"
	"github.com/sonr-io/sonr/pkg/ipns"
	"github.com/sonr-io/sonr/pkg/motor"
	mtu "github.com/sonr-io/sonr/testutil/motor"
	"github.com/sonr-io/sonr/third_party/types/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	st "github.com/sonr-io/sonr/x/schema/types"
)

type IPNSTestSuite struct {
	suite.Suite
	motorNode  motor.MotorNode
	testBucket bucket.BucketClient
	cid        string
}

func (suite *IPNSTestSuite) SetupSuite() {
	var err error

	suite.motorNode, err = motor.EmptyMotor(&mt.InitializeRequest{
		DeviceId: "test_device",
	}, common.DefaultCallback())
	if err != nil {
		fmt.Printf("Failed to setup test motor: %s\n", err)
	}

	err = mtu.SetupTestAddressWithKeys(suite.motorNode)
	if err != nil {
		fmt.Printf("Failed to setup test address: %s\n", err)
	}

	// create document
	// create schema
	createSchemaRequest := mt.CreateSchemaRequest{
		Label: "TestUser",
		Fields: map[string]*st.SchemaFieldKind{
			"email": {
				Kind: st.Kind_STRING,
			},
			"firstName": {
				Kind: st.Kind_STRING,
			},
			"age": {
				Kind: st.Kind_INT,
			},
		},
	}

	resp, err := suite.motorNode.CreateSchema(createSchemaRequest)

	// query WhatIs so it's cached
	_, err = suite.motorNode.QueryWhatIsByDid(resp.WhatIs.Did)

	// upload object
	builder, err := suite.motorNode.NewDocumentBuilder(resp.WhatIs.Did)

	builder.SetLabel("Player 1")
	builder.Set("email", "test_email")
	builder.Set("firstName", "test_name")
	builder.Set("age", 10)

	_, err = builder.Build()
	if err != nil {
		fmt.Printf("Failed to build document: %s\n", err)
	}

	result, err := builder.Upload()
	if "Player 1" != result.Document.Label {
		fmt.Println("Failed to upload document")
	}

	suite.cid = result.GetCid()
}

func Test_IPNS_Suite(t *testing.T) {
	suite.Run(t, new(IPNSTestSuite))
}

func (suite *IPNSTestSuite) Test_IPNS() {
	shell := shell.NewLocalShell()
	suite.T().Run("Should create ipns record", func(t *testing.T) {
		time_stamp := fmt.Sprintf("%d", time.Now().Unix())

		out_path := filepath.Join(os.TempDir(), time_stamp+".txt")

		defer os.Remove(out_path)
		priv, err := rsa.GenerateKey(rand.Reader, 2048)
		assert.NoError(t, err)
		rec, err := ipns.New(priv)
		assert.NoError(t, err)
		rec.Builder.SetCid(suite.cid)
		err = rec.CreateRecord()
		assert.NoError(t, err)
		srv := rec.Builder.BuildService()
		assert.NotNil(t, srv)
		fmt.Print(srv.ID)
		id, err := ipns.Publish(shell, rec)
		assert.NoError(t, err)
		str, err := ipns.Resolve(shell, id)
		assert.NoError(t, err)
		assert.NotNil(t, str)
		err = shell.Get(str, out_path)
		assert.NoError(t, err)
		buf, err := os.ReadFile(out_path)
		assert.NoError(t, err)
		fmt.Print(string(buf))
	})
}
