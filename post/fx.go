package post

import (
	"go.uber.org/fx"
	"main/post/Bloc"
	"main/post/handler"
	"main/post/query"
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
