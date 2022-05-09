package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login
func Login(w http.ResponseWriter, r *http.Request) {
	request, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(request, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repo := repositories.NewUserReposutory(db)
	userDB, erro := repo.FindByEmail(user.Email)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = auth.CheckPassword(userDB.Senha, user.Senha); erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := auth.GenerateToken(userDB.ID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	w.Write([]byte(token))
}
