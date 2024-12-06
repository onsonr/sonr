package config

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

// LoadFromBytes loads the environment from the given bytes
func LoadFromBytes(data []byte) (Env, error) {
	text := string(data)
	return LoadFromString(text)
}

// LoadFromString loads the environment from the given string
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

// LoadFromURL loads the environment from the given URL
func LoadFromURL(url string) (Env, error) {
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
	ret, err := Load(context.Background(), evaluator, pkl.UriSource(url))
	return ret, err
}
