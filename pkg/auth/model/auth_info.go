package model

import (
	"pawpawchat/generated/proto/auth"
	"time"
)

type AuthInfo struct {
	UserID    uint64    `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Email     string    `gorm:"type:varchar(64);unique;uniqueIndex;not null" json:"email"`
	HashPass  string    `gorm:"type:text;not null" json:"-"`
	Password  string    `json:"password"`
	LastLogin time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"last_login"`
}

func NewAuthInfo(in any) *AuthInfo {
	switch input := in.(type) {
	case *auth.SignUpRequest:
		return &AuthInfo{
			Email:    input.GetEmail(),
			Password: input.GetPassword(),
		}
	}

	return nil
}

// func (u *AuthInfo) UnmarshalJSON(data []byte) error {
// 	// declare structure to avoid recursive overloaded method
// 	var alias struct {
// 		UserID    uint64    `json:"user_id"`
// 		Email     string    `json:"email"`
// 		Password  string    `json:"password"`
// 		HashPass  string    `json:"hash_pass"`
// 		LastLogin time.Time `json:"last_login"`
// 	}

// 	if err := json.Unmarshal(data, &alias); err != nil {
// 		return err
// 	}

// 	u.UserID = alias.UserID
// 	u.Email = alias.Email
// 	u.Password = alias.Password
// 	u.HashPass = alias.HashPass
// 	u.LastLogin = alias.LastLogin

// 	return nil
// }

// func (u *AuthInfo) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(struct {
// 		UserID    uint64    `db:"user_id" json:"user_id"`
// 		Email     string    `db:"email" json:"email"`
// 		HashPass  string    `db:"hash_pass" json:"hash_pass"`
// 		LastLogin time.Time `db:"last_login" json:"last_login"`
// 	}{
// 		UserID:    u.UserID,
// 		Email:     u.Email,
// 		HashPass:  u.HashPass,
// 		LastLogin: u.LastLogin,
// 	})
// }
