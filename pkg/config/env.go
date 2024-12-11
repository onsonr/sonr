package config

import (
	"context"

	"github.com/apple/pkl-go/pkl"
	hwayconfig "github.com/onsonr/sonr/pkg/config/hway"
)

// LoadFromBytes loads the environment from the given bytes
func LoadHwayFromBytes(data []byte) (hwayconfig.Hway, error) {
	text := string(data)
	return LoadHwayFromString(text)
}

// LoadFromString loads the environment from the given string
func LoadHwayFromString(text string) (hwayconfig.Hway, error) {
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
	ret, err := hwayconfig.Load(context.Background(), evaluator, pkl.TextSource(text))
	return ret, err
}

// LoadFromURL loads the environment from the given URL
func LoadFromURL(url string) (hwayconfig.Hway, error) {
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
	ret, err := hwayconfig.Load(context.Background(), evaluator, pkl.UriSource(url))
	return ret, err
}
