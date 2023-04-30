package biz

import (
	"context"
	v1 "github.com/go-kratos/kratos-layout/api/helloworld/v1"
	"github.com/go-kratos/kratos-layout/internal/dao"
	"github.com/go-kratos/kratos/v2/log"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type GreeterBiz struct {
	log  *log.Helper
	repo dao.GreeterRepo
}

func NewGreeterBiz(repo dao.GreeterRepo, logger log.Logger) *GreeterBiz {
	return &GreeterBiz{log: log.NewHelper(logger), repo: repo}
}

func (biz *GreeterBiz) CreateGreeter(ctx context.Context, g *dao.Greeter) (*dao.Greeter, error) {
	biz.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return biz.repo.Save(ctx, g)
}
