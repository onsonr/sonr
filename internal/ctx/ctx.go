package ctx

import (
	"github.com/bool64/cache"
	"google.golang.org/grpc/metadata"
)

const (
	kMetaKeySession = "sonr-session-id"
)

var sessionCache *cache.FailoverOf[metadata.MD]
