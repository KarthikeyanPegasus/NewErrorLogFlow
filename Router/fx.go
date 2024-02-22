package Router

import (
	"go.uber.org/fx"
)

var Handler = fx.Options(
	fx.Provide(
		NewHandler,
	),
)
