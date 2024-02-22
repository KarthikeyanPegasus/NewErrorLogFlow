package sampleDatabase

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewConnectionServer),
)
