package config

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

func LoadFromBytes(data []byte) (Env, error) {
	text := string(data)
	return LoadFromString(text)
}

func LoadFromString(text string) (Env, error) {
	evaluator, err := pkl.NewEvaluator(context.Background(), pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err := Load(context.Background(), evaluator, pkl.TextSource(text))
	return ret, err
}
