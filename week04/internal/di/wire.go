// +build wireinject

package di

import (
	"week04/internal/biz"
	"week04/internal/dao"
	"week04/internal/server/grpc"
	"week04/internal/server/http"
	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, biz.Provider, http.New,  grpc.New, NewApp))
}

