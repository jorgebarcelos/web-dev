package controllers

import (
	"api/src/Db"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, read *http.Request){
	bodyRequest, err := io.ReadAll(read.Body)

	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, err)
		return
	}

	Db, err := db.Connect()

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
	}

	defer Db.Close()

	repository := repositories.NewUsersRepo(Db)
	user.ID, err = repository.Create(user)

	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)

}

func RetrieveUsers(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando usuários"))
}

func RetrieveUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando usuário"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizando usuário"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Excluindo usuário"))
}