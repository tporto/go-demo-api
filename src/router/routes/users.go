package routes

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:                "/users",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CreateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/users",
		Metodo:             http.MethodGet,
		Funcao:             controllers.FindUsers,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/users/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/users/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/users/{id}/seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.FollowUser,
		RequerAutenticacao: true,
	},
	{
		URI:                "/users/{id}/parar-seguir",
		Metodo:             http.MethodPost,
		Funcao:             controllers.StopFollowUser,
		RequerAutenticacao: true,
	},
}
