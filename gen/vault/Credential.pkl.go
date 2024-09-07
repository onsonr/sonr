// Code generated from Pkl module `vault`. DO NOT EDIT.
package vault

import "github.com/apple/pkl-go/pkl"

type Credential interface {
	Model

	GetId() uint

	GetSubject() string

	GetController() string

	GetAttestationType() string

	GetOrigin() string

	GetCredentialId() *pkl.Object

	GetPublicKey() *pkl.Object

	GetTransport() string

	GetSignCount() uint

	GetUserPresent() bool

	GetUserVerified() bool

	GetBackupEligible() bool

	GetBackupState() bool

	GetCloneWarning() bool

	GetCreatedAt() *string

	GetUpdatedAt() *string
}

var _ Credential = (*CredentialImpl)(nil)

type CredentialImpl struct {
	Table string `pkl:"table"`

	Id uint `pkl:"id" gorm:"primaryKey" json:"id,omitempty"`

	Subject string `pkl:"subject" json:"subject,omitempty"`

	Controller string `pkl:"controller" json:"controller,omitempty"`

	AttestationType string `pkl:"attestationType" json:"attestationType,omitempty"`

	Origin string `pkl:"origin" json:"origin,omitempty"`

	CredentialId *pkl.Object `pkl:"credentialId" json:"credentialId,omitempty"`

	PublicKey *pkl.Object `pkl:"publicKey" json:"publicKey,omitempty"`

	Transport string `pkl:"transport" json:"transport,omitempty"`

	SignCount uint `pkl:"signCount" json:"signCount,omitempty"`

	UserPresent bool `pkl:"userPresent" json:"userPresent,omitempty"`

	UserVerified bool `pkl:"userVerified" json:"userVerified,omitempty"`

	BackupEligible bool `pkl:"backupEligible" json:"backupEligible,omitempty"`

	BackupState bool `pkl:"backupState" json:"backupState,omitempty"`

	CloneWarning bool `pkl:"cloneWarning" json:"cloneWarning,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`

	UpdatedAt *string `pkl:"updatedAt" json:"updatedAt,omitempty"`
}

func (rcv *CredentialImpl) GetTable() string {
	return rcv.Table
}

func (rcv *CredentialImpl) GetId() uint {
	return rcv.Id
}

func (rcv *CredentialImpl) GetSubject() string {
	return rcv.Subject
}

func (rcv *CredentialImpl) GetController() string {
	return rcv.Controller
}

func (rcv *CredentialImpl) GetAttestationType() string {
	return rcv.AttestationType
}

func (rcv *CredentialImpl) GetOrigin() string {
	return rcv.Origin
}

func (rcv *CredentialImpl) GetCredentialId() *pkl.Object {
	return rcv.CredentialId
}

func (rcv *CredentialImpl) GetPublicKey() *pkl.Object {
	return rcv.PublicKey
}

func (rcv *CredentialImpl) GetTransport() string {
	return rcv.Transport
}

func (rcv *CredentialImpl) GetSignCount() uint {
	return rcv.SignCount
}

func (rcv *CredentialImpl) GetUserPresent() bool {
	return rcv.UserPresent
}

func (rcv *CredentialImpl) GetUserVerified() bool {
	return rcv.UserVerified
}

func (rcv *CredentialImpl) GetBackupEligible() bool {
	return rcv.BackupEligible
}

func (rcv *CredentialImpl) GetBackupState() bool {
	return rcv.BackupState
}

func (rcv *CredentialImpl) GetCloneWarning() bool {
	return rcv.CloneWarning
}

func (rcv *CredentialImpl) GetCreatedAt() *string {
	return rcv.CreatedAt
}

func (rcv *CredentialImpl) GetUpdatedAt() *string {
	return rcv.UpdatedAt
}
