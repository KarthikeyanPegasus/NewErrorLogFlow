package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/justinas/alice"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"main/Router"
	router "main/Router"
	database "main/database"
	logger "main/errorLogger"
	"main/middlewares"
	"main/post"
	"main/users"
	"net/http"
	"os"
)

var module = fx.Options(
	fx.Provide(
		ProvideDatabase,
		ProvideLogger,
		fx.Annotated{
			Name:   "http",
			Target: NewHttpServer,
		},
		DefaultRouter,
		NewChain,
	),
	middlewares.Module,
	logger.Module,
	users.Module,
	post.Module,
	Router.Handler,
	fx.Invoke(RunServer),
)

func ProvideDatabase() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	db, err := database.NewConnectionServer(&sql.DB{}).Create(dbHost, dbPort, dbUser, dbPassword, dbName)
	if err != nil {
		fmt.Println("Error while creating database connection: ", err)
		return nil
	}
	return db
}

func ProvideLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	config.DisableStacktrace = true
	zapper, _ := config.Build()
	return zapper
}

func NewHttpServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: handler,
	}
}

func RunServer(in struct {
	fx.In
	Lc         fx.Lifecycle
	HttpServer *http.Server `name:"http"`
}) {

	in.Lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				fmt.Println("Server is running on port 8080")
				in.HttpServer.ListenAndServe()
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return in.HttpServer.Shutdown(ctx)
		},
	})

}

func NewChain(in struct {
	fx.In
	Middleware *middlewares.MiddlewareLogger
}) *alice.Chain {
	c := alice.New(
		in.Middleware.CentralisedLogger(),
	)

	return &c
}

func DefaultRouter() router.Router {
	return chi.NewRouter()
}
