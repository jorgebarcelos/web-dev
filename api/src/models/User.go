package models

import "time"

type User struct {
	ID        uint64    `json:"id, omitempty"`
	UserName      string    `json:"user_name, omitempty"`
	Nick      string    `json:"nick, omitempty"`
	Email     string    `json:"emai, omitempty"`
	UserPassword  string    `json:"user_password, omitempty"`
	CreatedAt time.Time `json:"CreatedAt, omitempty"`
}
