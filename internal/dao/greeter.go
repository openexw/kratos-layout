package dao

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type (
	Greeter struct {
		Id         int64     `db:"id"`
		Number     string    `db:"number"`   // 学号
		Name       string    `db:"name"`     // 用户名称
		Password   string    `db:"password"` // 用户密码
		Gender     string    `db:"gender"`   // 男｜女｜未公开
		Hello      string    `db:"hello"`    // 男｜女｜未公开
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}

	GreeterRepo interface {
		Save(context.Context, *Greeter) (*Greeter, error)
		Update(context.Context, *Greeter) (*Greeter, error)
		FindByID(context.Context, int64) (*Greeter, error)
		ListByHello(context.Context, string) ([]*Greeter, error)
		ListAll(context.Context) ([]*Greeter, error)
	}

	greeterRepo struct {
		data *Dao
		log  *log.Helper
	}
)

func (*Greeter) TableName() string {
	return "greeters"
}

// NewGreeterRepo .
func NewGreeterRepo(data *Dao, logger log.Logger) GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *Greeter) (*Greeter, error) {
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *Greeter) (*Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*Greeter, error) {
	return nil, nil
}
