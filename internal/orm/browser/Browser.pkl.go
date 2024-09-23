// Code generated from Pkl module `browser`. DO NOT EDIT.
package browser

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Browser struct {
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Browser
func LoadFromPath(ctx context.Context, path string) (ret *Browser, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Browser
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Browser, error) {
	var ret Browser
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
