// Code generated from Pkl module `orm`. DO NOT EDIT.
package orm

type Credential struct {
	Id uint `pkl:"id" gorm:"primaryKey,autoIncrement" json:"id,omitempty" query:"id"`

	Subject string `pkl:"subject" json:"subject,omitempty" param:"subject"`

	Controller string `pkl:"controller" json:"controller,omitempty" param:"controller"`

	AttestationType string `pkl:"attestationType" json:"attestationType,omitempty" param:"attestationType"`

	Origin string `pkl:"origin" json:"origin,omitempty" param:"origin"`

	CredentialId string `pkl:"credentialId" json:"credentialId,omitempty" param:"credentialId"`

	PublicKey string `pkl:"publicKey" json:"publicKey,omitempty" param:"publicKey"`

	Transport string `pkl:"transport" json:"transport,omitempty" param:"transport"`

	SignCount uint `pkl:"signCount" json:"signCount,omitempty" param:"signCount"`

	UserPresent bool `pkl:"userPresent" json:"userPresent,omitempty" param:"userPresent"`

	UserVerified bool `pkl:"userVerified" json:"userVerified,omitempty" param:"userVerified"`

	BackupEligible bool `pkl:"backupEligible" json:"backupEligible,omitempty" param:"backupEligible"`

	BackupState bool `pkl:"backupState" json:"backupState,omitempty" param:"backupState"`

	CloneWarning bool `pkl:"cloneWarning" json:"cloneWarning,omitempty" param:"cloneWarning"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty" param:"createdAt"`

	UpdatedAt *string `pkl:"updatedAt" json:"updatedAt,omitempty" param:"updatedAt"`
}
