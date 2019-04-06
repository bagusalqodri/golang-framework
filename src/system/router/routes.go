package router

import (
	"learning-golang/pkg/type/routes"
	HomeHandler "learning-golang/src/controllers/home"
	"net/http"

	"github.com/go-xorm/xorm"
)

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func GetRoutes(db *xorm.Engine) routes.Routes {

	HomeHandler.Init(db)

	return routes.Routes{
		routes.Route{"Home", "GET", "/", HomeHandler.Index},
	}
}
