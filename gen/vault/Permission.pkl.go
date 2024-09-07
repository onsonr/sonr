// Code generated from Pkl module `vault`. DO NOT EDIT.
package vault

type Permission interface {
	Model

	GetId() uint

	GetServiceId() string

	GetGrants() string

	GetScopes() string

	GetCreatedAt() *string

	GetUpdatedAt() *string
}

var _ Permission = (*PermissionImpl)(nil)

type PermissionImpl struct {
	Table string `pkl:"table"`

	Id uint `pkl:"id" gorm:"primaryKey" json:"id,omitempty"`

	ServiceId string `pkl:"serviceId" json:"serviceId,omitempty"`

	Grants string `pkl:"grants" json:"grants,omitempty"`

	Scopes string `pkl:"scopes" json:"scopes,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`

	UpdatedAt *string `pkl:"updatedAt" json:"updatedAt,omitempty"`
}

func (rcv *PermissionImpl) GetTable() string {
	return rcv.Table
}

func (rcv *PermissionImpl) GetId() uint {
	return rcv.Id
}

func (rcv *PermissionImpl) GetServiceId() string {
	return rcv.ServiceId
}

func (rcv *PermissionImpl) GetGrants() string {
	return rcv.Grants
}

func (rcv *PermissionImpl) GetScopes() string {
	return rcv.Scopes
}

func (rcv *PermissionImpl) GetCreatedAt() *string {
	return rcv.CreatedAt
}

func (rcv *PermissionImpl) GetUpdatedAt() *string {
	return rcv.UpdatedAt
}
