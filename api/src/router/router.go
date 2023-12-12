package router

import "github.com/gorilla/mux"

// return configurated routers
func Generate() *mux.Router{
	return mux.NewRouter()
}