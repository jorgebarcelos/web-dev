package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

func (repositorie usersRepo) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, err := repositorie.db.Query(
		"select id, user_name, nick, email, created_at from users where user_name LIKE ? or nick LIKE ?",
		nameOrNick, nameOrNick)

	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID, 
			&user.UserName, 
			&user.Nick, 
			&user.Email, 
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}