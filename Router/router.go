package Router

import (
	"github.com/justinas/alice"
	"go.uber.org/fx"
	"net/http"
)

type Router interface {
	http.Handler

	Handle(pattern string, h http.Handler)
	HandleFunc(pattern string, h http.HandlerFunc)
	Delete(pattern string, h http.HandlerFunc)
	Get(pattern string, h http.HandlerFunc)
	Patch(pattern string, h http.HandlerFunc)
	Post(pattern string, h http.HandlerFunc)
	Put(pattern string, h http.HandlerFunc)
	Method(method, pattern string, h http.Handler)
}

type RouteFunc func(Router)

func NewHandler(in struct {
	fx.In
	Routes []RouteFunc `group:"routes"`
	Router Router
	Chain  *alice.Chain
}) (http.Handler, error) {
	for _, r := range in.Routes {
		r(in.Router)
	}

	if in.Chain == nil {
		return in.Router, nil
	}

	return in.Chain.Then(in.Router), nil
}
