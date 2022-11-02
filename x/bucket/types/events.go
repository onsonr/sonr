package types

const (
	EventTypeCreateWhereIs    = "createWhereIs"
	EventTypeGenerateBucket   = "generateBucket"
	EventTypeDeactivateBucket = "deactivateBucket"
	EventTypeUpdateBucket     = "updateBucket"
	EventTypeBurnBucket       = "burnBucket"

	AttributeValueCategory = ModuleName
	AttributeKeyTxType     = "txType"

	AttributeKeyCreator   = "bucketCreator"
	AttributeKeyServiceId = "serviceId"
	AttributeKeyLabel     = "label"
)
