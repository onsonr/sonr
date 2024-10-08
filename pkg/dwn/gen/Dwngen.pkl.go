// Code generated from Pkl module `dwngen`. DO NOT EDIT.
package gen

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type Dwngen struct {
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Dwngen
func LoadFromPath(ctx context.Context, path string) (ret *Dwngen, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Dwngen
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*Dwngen, error) {
	var ret Dwngen
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
