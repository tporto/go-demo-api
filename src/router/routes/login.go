package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRoute = Rota{
	URI:                "/login",
	Metodo:             http.MethodPost,
	Funcao:             controllers.Login,
	RequerAutenticacao: false,
}
