package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	UserName      string    `json:"user_name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	UserPassword  string    `json:"user_password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

//call Validate  and Format methods
func (user * User) Ready(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	user.format()
	return nil
}

// Verifies if are empty fields
func (user *User) validate (stage string) error {
	if user.UserName == "" {
		return errors.New("o campo 'nome' não pode estar vazio")
	}

	if user.Nick == "" {
		return errors.New("o campo 'apelido' não pode estar vazio")
	}

	if user.Email == "" {
		return errors.New("o campo 'email' não pode estar vazio")
	}

	if stage  == "register" && user.UserPassword == "" {
		return errors.New("o campo 'senha' não pode estar vazio")
	}

	return nil
}

// Remove empty spaces in extremities
func(user *User) format() {
	user.UserName = strings.TrimSpace(user.UserName)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}