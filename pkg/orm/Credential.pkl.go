// Code generated from Pkl module `orm`. DO NOT EDIT.
package orm

type Credential struct {
	Id string `pkl:"id" json:"id,omitempty" query:"id"`

	Subject string `pkl:"subject" json:"subject,omitempty"`

	Controller string `pkl:"controller" json:"controller,omitempty"`

	AttestationType string `pkl:"attestationType" json:"attestationType,omitempty"`

	Origin string `pkl:"origin" json:"origin,omitempty"`

	Label *string `pkl:"label" json:"label,omitempty"`

	DeviceId *string `pkl:"deviceId" json:"deviceId,omitempty"`

	CredentialId string `pkl:"credentialId" json:"credentialId,omitempty"`

	PublicKey string `pkl:"publicKey" json:"publicKey,omitempty"`

	Transport []string `pkl:"transport" json:"transport,omitempty"`

	SignCount uint `pkl:"signCount" json:"signCount,omitempty"`

	UserPresent bool `pkl:"userPresent" json:"userPresent,omitempty"`

	UserVerified bool `pkl:"userVerified" json:"userVerified,omitempty"`

	BackupEligible bool `pkl:"backupEligible" json:"backupEligible,omitempty"`

	BackupState bool `pkl:"backupState" json:"backupState,omitempty"`

	CloneWarning bool `pkl:"cloneWarning" json:"cloneWarning,omitempty"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty"`

	UpdatedAt *string `pkl:"updatedAt" json:"updatedAt,omitempty"`
}
