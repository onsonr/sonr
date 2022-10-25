package bucket_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/sonr-io/sonr/internal/bucket"
	"github.com/sonr-io/sonr/pkg/motor"
	mtu "github.com/sonr-io/sonr/testutil/motor"
	"github.com/sonr-io/sonr/third_party/types/common"
	mt "github.com/sonr-io/sonr/third_party/types/motor/api/v1"
	"github.com/sonr-io/sonr/x/bucket/types"
	bt "github.com/sonr-io/sonr/x/bucket/types"
	st "github.com/sonr-io/sonr/x/schema/types"
	"github.com/stretchr/testify/suite"
)

type BucketTestSuite struct {
	suite.Suite
	motorNode  motor.MotorNode
	testBucket bucket.Bucket
	cidDoc1    string
	cidDoc2    string
}

func (suite *BucketTestSuite) SetupSuite() {
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

	// create 2 different test items and upload them
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
	
	suite.cidDoc1 = result.GetCid()

	builder.SetLabel("Player 2")
	builder.Set("email", "test_email_2")
	builder.Set("firstName", "test_name_2")
	builder.Set("age", 20)

	result, err = builder.Upload()
	if "Player 2" != result.Document.Label {
		fmt.Println("Failed to upload document")
	}
	
	suite.cidDoc2 = result.GetCid()

	// create bucket with content
	content := []*bt.BucketItem{
		{
			Name:      "test",
			Uri:       suite.cidDoc1,
			Timestamp: time.Now().Unix(),
			Type:      bt.ResourceIdentifier_CID,
		},
	}

	_, suite.testBucket, _ = suite.motorNode.CreateBucket(mt.CreateBucketRequest{
		Creator:    suite.motorNode.GetAddress(),
		Label:      "test bucket",
		Visibility: types.BucketVisibility_PUBLIC,
		Role:       types.BucketRole_USER,
		Content:    content,
	})
}

func Test_BucketTestSuite(t *testing.T) {
	suite.Run(t, new(BucketTestSuite))
}
