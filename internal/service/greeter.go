package service

import (
	"context"
	"github.com/go-kratos/kratos-layout/internal/dao"

	v1 "github.com/go-kratos/kratos-layout/api/helloworld/v1"
	"github.com/go-kratos/kratos-layout/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc *biz.GreeterBiz
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterBiz) *GreeterService {
	return &GreeterService{uc: uc}
}

// SayHello implements helloworld.GreeterServer.
func (s *GreeterService) SayHello(ctx context.Context, req *v1.SayHelloReq) (*v1.SayHelloResp, error) {
	g, err := s.uc.CreateGreeter(ctx, &dao.Greeter{Hello: req.Name})
	if err != nil {
		return nil, err
	}
	return &v1.SayHelloResp{Message: "Hello " + g.Hello}, nil
}
