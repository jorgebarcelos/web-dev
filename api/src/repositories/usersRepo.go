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

func (repositorie usersRepo) Create(user models.User) (uint64, error) {
	statement, err := repositorie.db.Prepare(
		"insert into users (user_name, nick, email, user_password) values(?, ?, ?, ?)",
	)

	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.UserName, user.Nick, user.Email, user.UserPassword)

	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastInsertID), nil
}