package users

import (
	"go.uber.org/fx"
	"main/users/Bloc"
	"main/users/handler"
	"main/users/query"
)

var Module = fx.Options(
	query.Module,
	Bloc.Modules,
	handler.Modules,
	fx.Provide(
		NewRouter,
		fx.Annotated{
			Group:  "routes",
			Target: Routers,
		},
	),
)
