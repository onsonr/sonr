package actor

import (
	"context"

	"github.com/asynkron/protoactor-go/actor"
)

var keyshareSystem = actor.NewActorSystem()

func GetController(ctx context.Context) (Controller, error) {
	return nil, nil
}
