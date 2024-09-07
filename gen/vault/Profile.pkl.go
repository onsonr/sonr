// Code generated from Pkl module `vault`. DO NOT EDIT.
package vault

type Profile interface {
	Model

	GetId() string

	GetSubject() string

	GetController() string

	GetOriginUri() *string

	GetPublicMetadata() *string

	GetPrivateMetadata() *string

	GetCreatedAt() *string

	GetUpdatedAt() *string
}

var _ Profile = (*ProfileImpl)(nil)

type ProfileImpl struct {
	Table string `pkl:"table"`

	Id string `pkl:"id" gorm:"primaryKey" json:"id,omitempty"`

	Subject string `pkl:"subject" json:"subject,omitempty"`

	Controller string `pkl:"controller" json:"controller,omitempty"`

	OriginUri *string `pkl:"originUri" json:"originUri,omitempty"`

	PublicMetadata *string `pkl:"publicMetadata" json:"publicMetadata,omitempty"`

	PrivateMetadata *string `pkl:"privateMetadata" json:"privateMetadata,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`

	UpdatedAt *string `pkl:"updatedAt" json:"updatedAt,omitempty"`
}

func (rcv *ProfileImpl) GetTable() string {
	return rcv.Table
}

func (rcv *ProfileImpl) GetId() string {
	return rcv.Id
}

func (rcv *ProfileImpl) GetSubject() string {
	return rcv.Subject
}

func (rcv *ProfileImpl) GetController() string {
	return rcv.Controller
}

func (rcv *ProfileImpl) GetOriginUri() *string {
	return rcv.OriginUri
}

func (rcv *ProfileImpl) GetPublicMetadata() *string {
	return rcv.PublicMetadata
}

func (rcv *ProfileImpl) GetPrivateMetadata() *string {
	return rcv.PrivateMetadata
}

func (rcv *ProfileImpl) GetCreatedAt() *string {
	return rcv.CreatedAt
}

func (rcv *ProfileImpl) GetUpdatedAt() *string {
	return rcv.UpdatedAt
}
