package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate vai retornar um router com as rotas configuradas
func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Configure(r)
}
