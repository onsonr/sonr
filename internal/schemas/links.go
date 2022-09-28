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
func (as *schemaImpl) LoadSubSchemas(ctx context.Context) error {
	var links []string = make([]string, 0)
	for _, f := range as.fields {
		if f.GetKind() == st.Kind_LINK {
			if f.FieldKind.LinkDid == "" {
				return errSchemaFieldsInvalid
			}
			links = append(links, f.FieldKind.LinkDid)
		}

		kind := f.FieldKind
		for kind != nil {
			if kind.Kind == st.Kind_LINK {
				if kind.LinkDid == "" {
					return errSchemaFieldsInvalid
				}
				links = append(links, kind.LinkDid)
			}

			kind = kind.ListKind
		}
	}

	for len(links) > 0 {
		key := links[len(links)-1]
		links = links[:len(links)-1]
		buf, err := as.store.Get(ctx, key)

		if err != nil {
			return err
		}

		def := &st.Schema{}
		err = def.Unmarshal(buf)

		if err != nil {
			return err
		}

		as.subSchemas[key] = def

		for _, sf := range def.Fields {
			if sf.FieldKind.Kind == st.Kind_LINK {
				if sf.FieldKind.LinkDid == "" {
					return errSchemaFieldsInvalid
				}
				links = append(links, sf.FieldKind.LinkDid)
			}
		}
	}

	return nil
}
