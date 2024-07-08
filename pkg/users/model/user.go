package model

import (
	"pawpawchat/generated/proto/users"
	"strconv"
	"time"

	"database/sql"
)

type User struct {
	ID         uint64         `db:"id" json:"id"`
	Username   string         `db:"username" json:"username"`
	FirstName  string         `db:"first_name" json:"first_name"`
	SecondName string         `db:"second_name" json:"second_name"`
	Online     bool           `db:"online" json:"online"`
	LastSeen   time.Time      `db:"last_seen" json:"last_seen"`
	Age        int            `db:"age" json:"age"`
	Birthday   sql.NullString `db:"birthday" json:"birthday,omitempty"`
	CreatedAt  time.Time      `db:"created_at" json:"created_at"`
	IsBlocked  bool           `db:"is_blocked" json:"is_blocked"`
}

func NewUser(in any) *User {
	switch input := in.(type) {
	case *users.CreateRequest:
		return &User{
			ID:         input.GetUserID(),
			FirstName:  input.GetFirstName(),
			SecondName: input.GetSecondName(),
			Username:   generateUsername(input.GetUserID()),
		}
	}
	return nil
}

// REFACTOR
func generateUsername(id uint64) string {
	return "123_" + strconv.Itoa(int(id))
}
