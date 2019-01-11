package router

import (
	"learning-golang/pkg/type/routes"
	v1SubRoutes "learning-golang/src/controllers/v1/router"

	"github.com/gorilla/mux"
)

type Router struct {
	Router *mux.Router
}

func NewRouter() (r Router) {

	r.Router = mux.NewRouter().StrictSlash(true)

	return
}

func (r *Router) Init() {
	r.Router.Use(Middleware)
	baseRoutes := GetRoutes()
	for _, route := range baseRoutes {
		r.Router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	v1SubRoutes := v1SubRoutes.GetRoutes()
	for name, pack := range v1SubRoutes {
		r.AttachSubRouterWithMiddelware(name, pack.Routes, pack.Middleware)
	}
}

func (r *Router) AttachSubRouterWithMiddelware(path string, subroutes routes.Routes, middleware mux.MiddlewareFunc) (SubRouter *mux.Router) {
	SubRouter = r.Router.PathPrefix(path).Subrouter()
	SubRouter.Use(middleware)

	for _, route := range subroutes {
		SubRouter.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return
}
