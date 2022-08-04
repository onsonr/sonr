package object

func (ao *objectImpl) assert(object map[string]interface{}) error {
	return ao.schema.VerifyObject(object)
}
