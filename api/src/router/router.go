package router

import (
	"api/src/router/routers"
	"github.com/gorilla/mux"
)

// return configurated routers
func Generate() *mux.Router{
	r := mux.NewRouter()
	return routers.Configuration(r)
}