package dao

import (
	"github.com/go-kratos/kratos-layout/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is dao providers.
var ProviderSet = wire.NewSet(NewDao, NewGreeterRepo)

// Dao .
type Dao struct {
	// TODO wrapped database client
}

// NewDao .
func NewDao(c *conf.Data, logger log.Logger) (*Dao, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the dao resources")
	}
	return &Dao{}, cleanup, nil
}
