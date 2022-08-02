package object

func (ao *AppObjectInternalImpl) assert(object map[string]interface{}) error {
	return ao.schema.VerifyObject(object)
}
