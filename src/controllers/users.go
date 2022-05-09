package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser
func CreateUser(w http.ResponseWriter, r *http.Request) {
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

	if erro = user.Prepare("insert"); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repo := repositories.NewUserReposutory(db)
	user.ID, erro = repo.Create(user)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// UpdateUser
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	userToken, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userToken {
		responses.Erro(w, http.StatusForbidden, errors.New("Usuário diferente do autenticado."))
		return
	}

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User

	if erro = json.Unmarshal(body, &user); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("edit"); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repo := repositories.NewUserReposutory(db)
	if erro = repo.Update(userID, user); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FindUsers
func FindUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repo := repositories.NewUserReposutory(db)
	users, erro := repo.Find(nameOrNick)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, users)
}

// GetUser
func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
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
	user, erro := repo.FindByID(userID)
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// DeleteUser
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
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
	if erro = repo.Delete(userID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// FollowUser
func FollowUser(w http.ResponseWriter, r *http.Request) {
	followID, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)
	userID, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followID == userID {
		responses.Erro(w, http.StatusForbidden, errors.New("Não pode seguir você mesmo."))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repositories.NewUserReposutory(db)
	if erro = repo.Follow(userID, followID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// StopFollowUser
func StopFollowUser(w http.ResponseWriter, r *http.Request) {
	followID, erro := auth.ExtractUserID(r)
	if erro != nil {
		responses.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(r)
	userID, erro := strconv.ParseUint(params["id"], 10, 64)
	if erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if followID == userID {
		responses.Erro(w, http.StatusForbidden, errors.New("Não é possível parar de seguir você mesmo."))
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repositories.NewUserReposutory(db)
	if erro = repo.StopFollow(userID, followID); erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
