package interface

import (
	"log"
	"runtime"

	"github.com/getlantern/systray"
	"github.com/gobuffalo/packr"
	sonr "github.com/sonr-io/core/bind"
	md "github.com/sonr-io/core/internal/models"
	win "github.com/sonr-io/core/pkg/window"

	//"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Interface struct {
	node       *sonr.Node
}
