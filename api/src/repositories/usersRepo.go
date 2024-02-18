package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type usersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *usersRepo {
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

func (repositorie usersRepo) SearchByID(ID uint64) (models.User, error) {
	lines, err := repositorie.db.Query(
		"select id, user_name, nick, email, created_at from users where id = ?",
		ID)

	if err != nil {
		return models.User{}, err
	}

	defer lines.Close()

	var user models.User

	if lines.Next() {
		if err = lines.Scan(
			&user.ID,
			&user.UserName,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (repositorie usersRepo) Update(ID uint64, user models.User) error {
	statement, err := repositorie.db.Prepare(
		"update users set user_name = ?, nick = ?, email = ? where id = ? ",
	)

	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(user.UserName, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (repositorie usersRepo) Delete(ID uint64) error {
	statement, err := repositorie.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err = statement.Exec(ID); err != nil{return err}

	return nil
}
