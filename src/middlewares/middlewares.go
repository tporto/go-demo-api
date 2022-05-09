package middlewares

import (
	"api/src/auth"
	"api/src/responses"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if erro := auth.ValidateToken(r); erro != nil {
			responses.Erro(w, http.StatusUnauthorized, erro)

			return
		}

		next(w, r)
	}
}
