// Code generated from Pkl module `common.types.ORM`. DO NOT EDIT.
package models

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type ORM struct {
	DbName string `pkl:"db_name"`

	DbVersion int `pkl:"db_version"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a ORM
func LoadFromPath(ctx context.Context, path string) (ret *ORM, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a ORM
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*ORM, error) {
	var ret ORM
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
