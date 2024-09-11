// Code generated from Pkl module `orm`. DO NOT EDIT.
package orm

type PublicKey struct {
	Id uint `pkl:"id" gorm:"primaryKey,autoIncrement" json:"id,omitempty" query:"id"`

	Role int `pkl:"role" json:"role,omitempty" param:"role"`

	Algorithm int `pkl:"algorithm" json:"algorithm,omitempty" param:"algorithm"`

	Encoding int `pkl:"encoding" json:"encoding,omitempty" param:"encoding"`

	Jwk string `pkl:"jwk" json:"jwk,omitempty" param:"jwk"`

	CreatedAt *string `pkl:"createdAt" json:"createdAt,omitempty" param:"createdAt"`
}
