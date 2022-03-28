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

package cl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCL(t *testing.T) {
	params := GetDefaultParamSizes()
	attrCount := NewAttrCount(5, 1, 0) // TODO: integrate this into GetDefaultParamSizes

	org, err := NewOrg(params, attrCount)
	if err != nil {
		t.Errorf("error when generating CL org: %v", err)
	}

	// to generate new testing keys
	//WriteGob("../../client/testdata/clPubKey.gob", org.Keys.Pub)
	//WriteGob("../../client/testdata/clSecKey.gob", org.Keys.Sec)

	masterSecret := org.Keys.Pub.GenerateUserMasterSecret()

	cred := NewRawCred(attrCount)
	_ = cred.AddStrAttr("Name", "Jack", true)
	_ = cred.AddStrAttr("Gender", "M", true)
	_ = cred.AddStrAttr("Graduated", "true", true)
	_ = cred.AddInt64Attr("DateMin", 22342345, true)
	_ = cred.AddInt64Attr("DateMax", 32342345, true)
	_ = cred.AddInt64Attr("Age", 25, false)

	credMgr, err := NewCredManager(params, org.Keys.Pub, masterSecret, cred)
	if err != nil {
		t.Errorf("error when creating a user: %v", err)
	}

	credManagerPath := "../client/testdata/credManager.gob"
	WriteGob(credManagerPath, credMgr)

	credIssueNonceOrg := org.GetCredIssueNonce()

	credReq, err := credMgr.GetCredRequest(credIssueNonceOrg)
	if err != nil {
		t.Errorf("error when generating credential request: %v", err)
	}

	res, err := org.IssueCred(credReq)
	if err != nil {
		t.Errorf("error when issuing credential: %v", err)
	}

	// Store record to db
	mockDb := NewMockRecordManager()
	if err := mockDb.Store(credReq.Nym, res.Record); err != nil {
		t.Errorf("error saving record to db: %v", err)
	}

	userVerified, err := credMgr.Verify(res.Cred, res.AProof)
	if err != nil {
		t.Errorf("error when verifying credential: %v", err)
	}
	assert.Equal(t, true, userVerified, "credential proof not valid")

	// Before updating a credential, create a new Org object (obtaining and updating
	// credential usually don't happen at the same time)
	org, err = NewOrgFromParams(params, org.Keys)
	if err != nil {
		t.Errorf("error when generating CL org: %v", err)
	}

	// create new CredManager (updating or proving usually does not happen at the same time
	// as issuing)
	ReadGob(credManagerPath, credMgr)

	// TODO: update to rawcred
	a, _ := cred.GetAttr("Name")
	_ = a.UpdateValue("John")
	credMgr.Update(cred)

	rec, err := mockDb.Load(credMgr.Nym)
	if err != nil {
		t.Errorf("error saving record to db: %v", err)
	}

	newKnownAttrs := cred.GetKnownVals()
	res1, err := org.UpdateCred(credMgr.Nym, rec, credReq.Nonce, newKnownAttrs)
	if err != nil {
		t.Errorf("error when updating credential: %v", err)
	}
	if err := mockDb.Store(credMgr.Nym, res1.Record); err != nil {
		t.Errorf("error saving record to db: %v", err)
	}

	userVerified, err = credMgr.Verify(res1.Cred, res1.AProof)
	if err != nil {
		t.Errorf("error when verifying updated credential: %v", err)
	}
	assert.Equal(t, true, userVerified, "credential update failed")

	// Some other organization which would like to verify the credential can instantiate org without sec key.
	// It only needs Pub key of the organization that issued a credential.
	org, err = NewOrgFromParams(params, org.Keys)
	if err != nil {
		t.Errorf("error when generating CL org: %v", err)
	}

	revealedKnownAttrsIndices := []int{0}         // reveal only the first known attribute
	revealedCommitmentsOfAttrsIndices := []int{0} // reveal only the commitment of the first attribute (of those of which only commitments are known)

	nonce := org.GetProveCredNonce()
	randCred, proof, err := credMgr.BuildProof(res1.Cred, revealedKnownAttrsIndices,
		revealedCommitmentsOfAttrsIndices, nonce)
	if err != nil {
		t.Errorf("error when building credential proof: %v", err)
	}

	revealedKnownAttrs, revealedCommitmentsOfAttrs := credMgr.FilterAttributes(revealedKnownAttrsIndices,
		revealedCommitmentsOfAttrsIndices)

	cVerified, err := org.ProveCred(randCred.A, proof, revealedKnownAttrsIndices,
		revealedCommitmentsOfAttrsIndices, revealedKnownAttrs, revealedCommitmentsOfAttrs)
	if err != nil {
		t.Errorf("error when verifying credential: %v", err)
	}

	assert.Equal(t, true, cVerified, "credential verification failed")
}
