package schemas

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	st "github.com/sonr-io/sonr/x/schema/types"
)

func (as *appSchemaInternalImpl) GetSchemaByCid(cid string) ([]*st.SchemaKindDefinition, error) {
	tmpPath := fmt.Sprintf("%s%s%d", os.TempDir(), cid, time.Now().Unix())
	err := as.shell.Get(cid, tmpPath)
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpPath)

	buf, err := os.ReadFile(tmpPath)

	if err != nil {
		return nil, err
	}
	content := make(map[string]interface{})
	if err = json.Unmarshal(buf, &content); err != nil {
		return nil, err
	}

	fieldsList := make([]*st.SchemaKindDefinition, len(content))
	for k, v := range content {
		fieldsList = append(fieldsList, &st.SchemaKindDefinition{
			Name:  k,
			Field: v.(st.SchemaKind),
		})
	}

	as.schemas[cid] = fieldsList
	return fieldsList, nil
}

func (as *appSchemaInternalImpl) GetWhatIs(creator string, did string) (*st.WhatIs, error) {
	res, err := as.client.QueryWhatIsByController(creator, did)
	if err != nil {
		return nil, err
	}
	as.whatIs[did] = res

	return res, nil
}
