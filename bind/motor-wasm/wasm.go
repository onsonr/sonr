package motor

import (
	"context"

	"github.com/sonr-io/sonr/pkg/config"
	"github.com/sonr-io/sonr/pkg/host"
)

func main() {
	config := config.DefaultConfig(config.Role_MOTOR)
	_, err := host.NewWasmHost(context.Background(), config)
	check(err)

}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
