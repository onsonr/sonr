package schemas

import (
	"context"

	st "github.com/sonr-io/sonr/x/schema/types"
)

/*
	Preprocessor for the top level schema to cache node paths for all sub schemas,
	bypassing link loader implementation due to lack of compatibilitty with our arch. Once support for json schemas are added we will no longer need to parse in this structure.
	but will loose the ability to reuse sub schemas in this fashion.
*/
func (as *schemaImpl) loadSubSchemas(ctx context.Context, fields []*st.SchemaKindDefinition) error {
	var links []string = make([]string, 0)
	for _, f := range fields {
		if f.LinkKind == st.LinkKind_SCHEMA {
			if f.Link == "" {
				return errSchemaFieldsInvalid
			}
			links = append(links, f.Link)
		}
	}

	for len(links) > 0 {
		key := links[len(links)-1]
		links = links[:len(links)-1]
		buf, err := as.store.Get(ctx, key)

		if err != nil {
			return err
		}

		var def st.SchemaDefinition
		err = def.Unmarshal(buf)

		if err != nil {
			return err
		}

		as.subSchemas[key] = &def

		for _, sf := range def.Fields {
			if sf.LinkKind == st.LinkKind_SCHEMA {
				if sf.Link == "" {
					return errSchemaFieldsInvalid
				}
				links = append(links, sf.Link)
			}
		}
	}

	return nil
}
