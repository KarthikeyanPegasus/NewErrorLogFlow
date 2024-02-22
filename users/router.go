package users

import (
	router "main/Router"
	handler "main/users/handler"
)

type Routes struct {
	handler *handler.Handler
}

func NewRouter(handler *handler.Handler) *Routes {
	return &Routes{handler: handler}
}

func Routers(rr *Routes) router.RouteFunc {
	return func(r router.Router) {
		r.Get("/users/{id}", rr.handler.Get)
		r.Post("/users", rr.handler.Create)
		r.Put("/users/{id}", rr.handler.Update)
		r.Delete("/users/{id}", rr.handler.Delete)
		r.Post("/users/upsert", rr.handler.Upsert)
	}
}
