package types

const (
	EventTypeCreateWhoIs     = "createWhoIs"
	EventTypeUpdateWhoIs     = "updateWhoIs"
	EventTypeDeactivateWhoIs = "deactivateWhoIs"

	AttributeValueCategory = ModuleName
	AttributeKeyTxType     = "txType"

	AttributeKeyCreator = "whoIsCreator"
	AttributeKeyDID     = "did"

	AttributeKeyOwner  = "whoIsOwner"
	AttributeKeySeller = "whoIsSeller"
	AttributeKeyAlias  = "aliasName"

	EventTypeBuyAlias      = "buyAlias"
	EventTypeSellAlias     = "sellAlias"
	EventTypeTransferAlias = "transferAlias"
)
