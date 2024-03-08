package post

import (
	router "main/Router"
	"main/post/handler"
)

type Routes struct {
	handler *handler.Handler
}

func NewRouter(handler *handler.Handler) *Routes {
	return &Routes{handler: handler}
}

func Routers(rr *Routes) router.RouteFunc {
	return func(r router.Router) {
		r.Get("/posts/{id}", rr.handler.Get)
		r.Post("/posts", rr.handler.Create)
		r.Put("/posts/{id}", rr.handler.Update)
		r.Delete("/posts/{id}", rr.handler.Delete)
		r.Post("/posts/delete", rr.handler.DeleteAllPostOfUser)
	}
}
