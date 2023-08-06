package sql

import "context"

var ifq *IceFireMySQL

type IceFireMySQL struct {
	ctx context.Context
}
