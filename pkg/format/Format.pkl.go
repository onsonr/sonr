// Code generated from Pkl module `format`. DO NOT EDIT.
package format

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Format struct {
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Format
func LoadFromPath(ctx context.Context, path string) (ret *Format, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Format
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Format, error) {
	var ret Format
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
