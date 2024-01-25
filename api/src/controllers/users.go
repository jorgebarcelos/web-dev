package controllers

import (
	"api/src/Db"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func CreateUser(write http.ResponseWriter, read *http.Request){
	bodyRequest, err := io.ReadAll(read.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User

	if err = json.Unmarshal(bodyRequest, &user); err != nil {
		log.Fatal(err)
	}

	Db, err := db.Connect()

	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUsersRepo(Db)
	userID, err := repository.Create(user)

	if err != nil {
		log.Fatal(err)
	}

	write.Write([]byte(fmt.Sprintf("Id inserido: %d", userID)))

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