package main

import (
	"fmt"
	"go.uber.org/fx"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

func run() error {

	app := fx.New(
		module,
	)
	if err := app.Err(); err != nil {
		return err
	}

	app.Run()

	return nil
}
