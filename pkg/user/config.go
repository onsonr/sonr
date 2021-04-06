import (
	"path/filepath"

	"google.golang.org/protobuf/proto"

	"github.com/libp2p/go-libp2p-core/crypto"

	"github.com/pkg/errors"
	md "github.com/sonr-io/core/pkg/models"
)

type UserConfig struct {
	connectivity *md.Connectivity

}
