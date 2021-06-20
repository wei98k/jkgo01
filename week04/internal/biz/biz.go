package biz

import (
	pb "week04/api/hello/v1"
	"week04/internal/dao"
	"context"
	"fmt"
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New, wire.Bind(new(pb.HelloServer), new(*Biz)))

// 自动化生成。
type Biz struct {
	ac  *paladin.Map
	dao dao.Dao
}

func New(d dao.Dao) (s *Biz, cf func(), err error) {
	s = &Biz{
		ac:  &paladin.TOML{},
		dao: d,
	}
	cf = s.Close
	err = paladin.Watch("application.toml", s.ac)
	return
}

// SayHello grpc func.
func (s *Biz) SayHello(ctx context.Context, req *pb.HelloReq) (reply *empty.Empty, err error) {
	reply = new(empty.Empty)

	fmt.Printf("hello %s", req.Name)
	return
}

func (s *Biz) Close() {}


