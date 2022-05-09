package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configure : coloca todas as rotas dentro do router
func Configure(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, loginRoute)

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI, middlewares.Authenticate(rota.Funcao)).Methods(rota.Metodo)
		} else {
			r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
		}

	}

	return r
}
