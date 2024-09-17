package service

import (
	"context"
	v1 "github.com/go-kratos/kratos-layout/gen/helloworld/v1"
	"github.com/go-kratos/kratos-layout/internal/biz"
	"github.com/go-kratos/kratos-layout/internal/dao"
)

// GreeterService is a greeter service.
type GreeterService struct {
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

//func (s *GreeterService) SayHello(ctx context.Context, req *connect.Request[v1.SayHelloReq]) (resp *connect.Response[v1.SayHelloResp], err error) {
//	s.uc.CreateGreeter()
//	name := req.Msg.GetName()
//	return connect.NewResponse(&v1.SayHelloResp{
//		Message: name,
//	}), nil
//}
