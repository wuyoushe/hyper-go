// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package wire

import (
	"github.com/google/wire"

	"hyper-go/tool\hyper/felton_blog/internal/config"
	"hyper-go/tool\hyper/felton_blog/internal/db"
	"hyper-go/tool\hyper/felton_blog/internal/server"
	"hyper-go/tool\hyper/felton_blog/internal/service"
)

var initProvider = wire.NewSet(config.NewConfig, db.NewDB, db.NewMC)
var svcProvider = wire.NewSet(service.NewHelperMap, service.New, db.NewCasbin)
var httpProvider = wire.NewSet(server.InitRouter, server.NewHttpServer)

//go:generate wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(
		initProvider,
		svcProvider,
		httpProvider,
		NewApp,
	))
}
