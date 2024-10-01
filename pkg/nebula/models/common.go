package models

import (
	"context"
	"errors"

	"github.com/apple/pkl-go/pkl"
)

var models *Models

func LoadFromString(ctx context.Context, s string) (err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err := Load(ctx, evaluator, pkl.TextSource(s))
	if err != nil {
		return err
	}
	models = ret
	return nil
}

func GetModels() (*Models, error) {
	if models == nil {
		return nil, errors.New("models not initialized")
	}
	return models, nil
}
