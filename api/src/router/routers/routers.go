package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// "Class" routes
type Router struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	AuthRequire bool
}


func Configuration(r *mux.Router) *mux.Router{
	routes := usersRouter

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}