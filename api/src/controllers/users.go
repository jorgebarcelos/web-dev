package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Criando usuário"))
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