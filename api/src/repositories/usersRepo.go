package repositories

import (
	"api/src/models"
	"database/sql"
)


type usersRepo struct {
	db * sql.DB
}

func NewUsersRepo(db * sql.DB) *usersRepo {
	return &usersRepo{db}
}

func (u usersRepo) Create(user models.User) (uint64, error) {
	return 0, nil
}